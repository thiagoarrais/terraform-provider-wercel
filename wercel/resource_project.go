package wercel

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
							ValidateDiagFunc: toValidateDiagFunc("type", validation.StringInSlice([]string{"gitlab"}, true)),
						},
						"project_url": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceProjectCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{}
	token := m.(string)

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	projectURL := d.Get("repo").([]interface{})[0].(map[string]interface{})["project_url"].(string)

	// TODO: what about gitlab subgroups?
	re, _ := regexp.Compile("([^/]+)/([^/]+)$")
	matches := re.FindStringSubmatch(projectURL)
	gitlabNamespace := matches[1]
	gitlabProjectName := matches[2]

	// TODO: should only run this if can't read project
	// WARN: undocumented endpoint POST /v6/projects
	reqBody, err := json.Marshal(map[string]interface{}{
		"name": name,
		"gitRepository": map[string]interface{}{
			"type":       "gitlab",
			"sourceless": true,
			"repo":       fmt.Sprintf("%s/%s", gitlabNamespace, gitlabProjectName),
		},
	})
	if err != nil {
		return diag.FromErr(err)
	}
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/v6/projects", "https://api.vercel.com"),
		bytes.NewReader(reqBody),
	)
	if err != nil {
		return diag.FromErr(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		if len(resp.Header["Content-Type"]) != 1 || resp.Header["Content-Type"][0] != "application/json; charset=utf-8" {
			return diag.Errorf("unknown error: %d %s", resp.StatusCode, resp.Status)
		}
		var errResponse map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&errResponse)
		if err != nil {
			return diag.FromErr(err)
		}
		errMessage := errResponse["error"].(map[string]interface{})["message"].(string)
		return diag.Errorf(errMessage)
	}

	var createProject map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&createProject)
	if err != nil {
		return diag.FromErr(err)
	}
	gitSource := createProject["link"].(map[string]interface{})
	gitSourceProjectID := gitSource["projectId"].(string)
	gitSourceProductionBranch := gitSource["productionBranch"].(string)
	// WARN: undocumented endpoint POST /v13/now/deployments
	reqBody, err = json.Marshal(map[string]interface{}{
		"name":   name,
		"target": "production",
		"source": "import",
		"gitSource": map[string]interface{}{
			"type":      "gitlab",
			"ref":       gitSourceProductionBranch,
			"projectId": gitSourceProjectID,
		},
	})
	if err != nil {
		return diag.FromErr(err)
	}
	req, err = http.NewRequest(
		"POST",
		fmt.Sprintf("%s/v13/now/deployments", "https://api.vercel.com"),
		bytes.NewReader(reqBody),
	)
	if err != nil {
		return diag.FromErr(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	resp, err = client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		if len(resp.Header["Content-Type"]) != 1 || resp.Header["Content-Type"][0] != "application/json; charset=utf-8" {
			return diag.Errorf("unknown error: %d %s", resp.StatusCode, resp.Status)
		}
		var errResponse map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&errResponse)
		if err != nil {
			return diag.FromErr(err)
		}
		errMessage := errResponse["error"].(map[string]interface{})["message"].(string)
		return diag.Errorf(errMessage)
	}

	d.SetId(createProject["id"].(string))

	return diags
}

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{}
	token := m.(string)

	var diags diag.Diagnostics

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/projects/%s", "https://api.vercel.com", d.Id()), nil)
	if err != nil {
		return diag.FromErr(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	var project map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&project)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", project["name"].(string)); err != nil {
		return diag.FromErr(err)
	}

	if project["link"] != nil {
		projectLink := project["link"].(map[string]interface{})
		repo := map[string]interface{}{
			"type":        projectLink["type"],
			"project_url": projectLink["projectUrl"],
		}
		d.Set("repo", []interface{}{repo})
	}

	d.SetId(project["id"].(string))

	return diags
}

func resourceProjectUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	if d.HasChange("repo") {
		client := &http.Client{}
		token := m.(string)

		// WARN: undocumented endpoint DELETE /v4/projects/:project_id/link
		req, err := http.NewRequest(
			"DELETE",
			fmt.Sprintf("%s/v4/projects/%s/link", "https://api.vercel.com", d.Id()),
			nil,
		)
		if err != nil {
			return diag.FromErr(err)
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

		resp, err := client.Do(req)
		if err != nil {
			return diag.FromErr(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			if len(resp.Header["Content-Type"]) != 1 || resp.Header["Content-Type"][0] != "application/json; charset=utf-8" {
				return diag.Errorf("unknown error: %d %s", resp.StatusCode, resp.Status)
			}
			var errResponse map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&errResponse)
			if err != nil {
				return diag.FromErr(err)
			}
			errMessage := errResponse["error"].(map[string]interface{})["message"].(string)
			return diag.Errorf(errMessage)
		}

		projectURL := d.Get("repo").([]interface{})[0].(map[string]interface{})["project_url"].(string)

		// TODO: what about gitlab subgroups?
		re, _ := regexp.Compile("([^/]+)/([^/]+)$")
		matches := re.FindStringSubmatch(projectURL)
		gitlabNamespace := matches[1]
		gitlabProjectName := matches[2]

		// WARN: undocumented endpoint POST /v4/projects/:project_id/link
		reqBody, err := json.Marshal(map[string]interface{}{
			"type": "gitlab",
			"repo": fmt.Sprintf("%s/%s", gitlabNamespace, gitlabProjectName),
		})
		if err != nil {
			return diag.FromErr(err)
		}
		req, err = http.NewRequest(
			"POST",
			fmt.Sprintf("%s/v4/projects/%s/link", "https://api.vercel.com", d.Id()),
			bytes.NewReader(reqBody),
		)
		if err != nil {
			return diag.FromErr(err)
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		req.Header.Add("Content-Type", "application/json; charset=utf-8")

		resp, err = client.Do(req)
		if err != nil {
			return diag.FromErr(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			if len(resp.Header["Content-Type"]) != 1 || resp.Header["Content-Type"][0] != "application/json; charset=utf-8" {
				return diag.Errorf("unknown error: %d %s", resp.StatusCode, resp.Status)
			}
			var errResponse map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&errResponse)
			if err != nil {
				return diag.FromErr(err)
			}
			errMessage := errResponse["error"].(map[string]interface{})["message"].(string)
			return diag.Errorf(errMessage)
		}

		var linkProject map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&linkProject)
		if err != nil {
			return diag.FromErr(err)
		}
		gitSource := linkProject["link"].(map[string]interface{})
		gitSourceProjectID := gitSource["projectId"].(string)
		gitSourceProductionBranch := gitSource["productionBranch"].(string)
		name := d.Get("name").(string)
		// WARN: undocumented endpoint POST /v13/now/deployments
		reqBody, err = json.Marshal(map[string]interface{}{
			"name":   name,
			"target": "production",
			"source": "import",
			"gitSource": map[string]interface{}{
				"type":      "gitlab",
				"ref":       gitSourceProductionBranch,
				"projectId": gitSourceProjectID,
			},
		})
		if err != nil {
			return diag.FromErr(err)
		}
		req, err = http.NewRequest(
			"POST",
			fmt.Sprintf("%s/v13/now/deployments", "https://api.vercel.com"),
			bytes.NewReader(reqBody),
		)
		if err != nil {
			return diag.FromErr(err)
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		req.Header.Add("Content-Type", "application/json; charset=utf-8")

		resp, err = client.Do(req)
		if err != nil {
			return diag.FromErr(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			if len(resp.Header["Content-Type"]) != 1 || resp.Header["Content-Type"][0] != "application/json; charset=utf-8" {
				return diag.Errorf("unknown error: %d %s", resp.StatusCode, resp.Status)
			}
			var errResponse map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&errResponse)
			if err != nil {
				return diag.FromErr(err)
			}
			errMessage := errResponse["error"].(map[string]interface{})["message"].(string)
			return diag.Errorf(errMessage)
		}

	}

	return diags
}

func resourceProjectDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{}
	token := m.(string)

	var diags diag.Diagnostics

	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf("%s/v1/projects/%s", "https://api.vercel.com", d.Id()),
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		if len(resp.Header["Content-Type"]) != 1 || resp.Header["Content-Type"][0] != "application/json; charset=utf-8" {
			return diag.Errorf("unknown error: %d %s", resp.StatusCode, resp.Status)
		}
		var errResponse map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&errResponse)
		if err != nil {
			return diag.FromErr(err)
		}
		errMessage := errResponse["error"].(map[string]interface{})["message"].(string)
		return diag.Errorf(errMessage)
	}

	d.SetId("")

	return diags
}
