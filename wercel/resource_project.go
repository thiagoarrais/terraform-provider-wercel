package wercel

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/thiagoarrais/terraform-provider-wercel/sdk"
)

func resourceProject() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		CreateContext: resourceProjectCreate,
		ReadContext:   resourceProjectRead,
		UpdateContext: resourceProjectUpdate,
		DeleteContext: resourceProjectDelete,
		Schema: map[string]*schema.Schema{
			"domains": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"repo": {
				Type:     schema.TypeList,
				Optional: true, // (?)
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							ValidateDiagFunc: toValidateDiagFunc("type", validation.StringInSlice([]string{"gitlab", "github"}, true)),
						},
						"project_url": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"root": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
		},
	}
}

func resourceProjectCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sdkClient := m.(config).sdkClient

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	repoType := d.Get("repo").([]interface{})[0].(map[string]interface{})["type"].(string)
	projectURL := d.Get("repo").([]interface{})[0].(map[string]interface{})["project_url"].(string)
	root := d.Get("root").(string)

	var rootDirectory *string
	if root != "" {
		rootDirectory = &root
	}

	// TODO: what about gitlab subgroups?
	re, _ := regexp.Compile("([^/]+)/([^/]+)$")
	matches := re.FindStringSubmatch(projectURL)
	gitlabNamespace := matches[1]
	gitlabProjectName := matches[2]

	tru := true

	// TODO: should only run this if can't read project
	project, _, err := sdkClient.ProjectsApi.CreateProject(ctx).
		ProjectCreation(sdk.ProjectCreation{
			Name: name,
			GitRepository: &sdk.GitRepositoryLink{
				Type:       repoType,
				Sourceless: &tru,
				Repo:       fmt.Sprintf("%s/%s", gitlabNamespace, gitlabProjectName),
			},
			RootDirectory: rootDirectory,
		}).
		WithUserCredentials(1).
		Execute()
	if err != nil {
		return diagFromSDKErr(err)
	}

	projectLink := project.GetLink()
	gitSource := sdk.NewDeploymentCreationGitSource()
	gitSource.SetType(repoType)
	gitSource.SetRef(projectLink.GetProductionBranch())
	gitSource.ProjectId = projectLink.ProjectId
	gitSource.RepoId = projectLink.RepoId
	deploymentCreation := sdk.NewDeploymentCreation(name)
	deploymentCreation.SetTarget("production")
	deploymentCreation.SetSource("import")
	deploymentCreation.SetGitSource(*gitSource)
	_, err = sdkClient.DeploymentsApi.CreateNewDeployment(ctx).
		DeploymentCreation(*deploymentCreation).
		Execute()
	if err != nil {
		return diagFromSDKErr(err)
	}

	if d.HasChange("domains") {
		old, new := d.GetChange("domains")
		err := syncDomains(ctx, project.GetId(), sdkClient, old, new)
		if err != nil {
			return diagFromSDKErr(err)
		}
	}

	d.SetId(project.GetId())

	return diags
}

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sdkClient := m.(config).sdkClient

	var diags diag.Diagnostics

	project, _, err := sdkClient.ProjectsApi.GetProjectById(ctx, d.Id()).Execute()
	if err != nil {
		return diagFromSDKErr(err)
	}

	if err := d.Set("name", project.GetName()); err != nil {
		return diag.FromErr(err)
	}

	if link, ok := project.GetLinkOk(); ok {
		repo := map[string]interface{}{
			"type":        link.GetType(),
			"project_url": link.GetProjectUrl(),
		}
		d.Set("repo", []interface{}{repo})
	}

	if aliases, ok := project.GetAliasOk(); ok {
		defaultDomain := project.GetName() + ".vercel.app"
		domains := []string{}
		for _, alias := range *aliases {
			if alias.GetDomain() != defaultDomain {
				domains = append(domains, alias.GetDomain())
			}
		}
		d.Set("domains", domains)
	}

	d.SetId(project.GetId())

	return diags
}

func resourceProjectUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	sdkClient := m.(config).sdkClient

	if d.HasChange("repo") {
		_, _, err := sdkClient.ProjectsApi.RemoveLinkByProjectId(ctx, d.Id()).Execute()
		if err != nil {
			return diagFromSDKErr(err)
		}

		repoType := d.Get("repo").([]interface{})[0].(map[string]interface{})["type"].(string)
		projectURL := d.Get("repo").([]interface{})[0].(map[string]interface{})["project_url"].(string)

		// TODO: what about gitlab subgroups?
		re, _ := regexp.Compile("([^/]+)/([^/]+)$")
		matches := re.FindStringSubmatch(projectURL)
		gitlabNamespace := matches[1]
		gitlabProjectName := matches[2]

		linkProject, _, err := sdkClient.ProjectsApi.CreateLinkByProjectId(ctx, d.Id()).
			GitRepositoryLink(sdk.GitRepositoryLink{
				Type: repoType,
				Repo: fmt.Sprintf("%s/%s", gitlabNamespace, gitlabProjectName),
			}).Execute()
		if err != nil {
			return diagFromSDKErr(err)
		}

		gitSource := linkProject.GetLink()
		gitSourceProductionBranch := gitSource.GetProductionBranch()
		name := d.Get("name").(string)
		gitRepositoryLink := sdk.NewDeploymentCreationGitSource()
		gitRepositoryLink.SetType(repoType)
		gitRepositoryLink.SetRef(gitSourceProductionBranch)
		gitRepositoryLink.ProjectId = gitSource.ProjectId
		gitRepositoryLink.RepoId = gitSource.RepoId
		deploymentCreation := sdk.NewDeploymentCreation(name)
		deploymentCreation.SetTarget("production")
		deploymentCreation.SetSource("import")
		deploymentCreation.SetGitSource(*gitRepositoryLink)
		_, err = sdkClient.DeploymentsApi.CreateNewDeployment(ctx).
			DeploymentCreation(*deploymentCreation).
			Execute()
		if err != nil {
			return diagFromSDKErr(err)
		}
	}

	if d.HasChange("domains") {
		old, new := d.GetChange("domains")
		err := syncDomains(ctx, d.Id(), sdkClient, old, new)
		if err != nil {
			return diagFromSDKErr(err)
		}
	}

	return diags
}

func syncDomains(ctx context.Context, projectID string, sdkClient *sdk.APIClient, old interface{}, new interface{}) error {
	oldSet := old.(*schema.Set)
	newSet := new.(*schema.Set)
	removedDomains := oldSet.Difference(newSet)
	addedDomains := newSet.Difference(oldSet)

	for _, removedDomain := range removedDomains.List() {
		_, err := sdkClient.ProjectsApi.RemoveAliasFromProject(ctx, projectID).Domain(removedDomain.(string)).Execute()
		if err != nil {
			return err
		}
	}

	for _, addedDomain := range addedDomains.List() {
		_, err := sdkClient.ProjectsApi.AddAliasToProject(ctx, projectID).
			AliasInclusion(sdk.AliasInclusion{
				Domain: addedDomain.(string),
			}).
			Execute()
		if err != nil {
			return err
		}
	}

	return nil
}

func resourceProjectDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sdkClient := m.(config).sdkClient

	var diags diag.Diagnostics

	_, err := sdkClient.ProjectsApi.RemoveProjectById(ctx, d.Id()).
		Execute()
	if err != nil {
		return diagFromSDKErr(err)
	}

	d.SetId("")

	return diags
}
