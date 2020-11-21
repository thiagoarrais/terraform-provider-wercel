package wercel

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/thiagoarrais/terraform-provider-wercel/sdk"
)

func toValidateDiagFunc(fieldName string, validate schema.SchemaValidateFunc) schema.SchemaValidateDiagFunc {
	return func(input interface{}, path cty.Path) diag.Diagnostics {
		var diags diag.Diagnostics
		_, errs := validate(input, fieldName)
		for _, err := range errs {
			diags = append(diags, diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       err.Error(),
				AttributePath: path,
			})
		}
		return diags
	}
}

func stringIgnoreCase(k string, old string, new string, d *schema.ResourceData) bool {
	return strings.EqualFold(old, new)
}

func resourceEnvvar() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		CreateContext: resourceEnvvarCreate,
		ReadContext:   resourceEnvvarRead,
		DeleteContext: resourceEnvvarDelete,
		CustomizeDiff: resourceEnvvarDiff,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"target": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				Default:          "production",
				DiffSuppressFunc: stringIgnoreCase,
				ValidateDiagFunc: toValidateDiagFunc("target", validation.StringInSlice([]string{"development", "preview", "production"}, true)),
			},
			"secret_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secret_uid": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceEnvvarCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	token := m.(string)

	var diags diag.Diagnostics

	value := d.Get("value").(string)
	secretName := uuid.New().String()

	conf := sdk.NewConfiguration()
	conf.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", token))
	sdkClient := sdk.NewAPIClient(conf)
	secret, _, err := sdkClient.SecretsApi.CreateNewSecret(ctx).
		SecretCreation(sdk.SecretCreation{
			Name:  secretName,
			Value: value,
		}).
		Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	secretUID := secret.GetUid()
	projectID := d.Get("project_id").(string)
	key := d.Get("key").(string)
	target := d.Get("target").(string)

	_, _, err = sdkClient.ProjectsApi.CreateProjectEnvironmentVariable(ctx, projectID).
		EnvironmentVariableCreation(sdk.EnvironmentVariableCreation{
			Type:   "secret",
			Key:    key,
			Value:  secretUID,
			Target: []string{strings.ToLower(target)},
		}).
		Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%s/%s", projectID, key))
	d.Set("secret_name", secretName)
	d.Set("secret_uid", secretUID)

	return diags
}

func resourceEnvvarRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// no-op because secret validation happens on resourceEnvvarDiff
	var diags diag.Diagnostics
	return diags
}

func resourceEnvvarDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{}
	token := m.(string)

	var diags diag.Diagnostics

	projectID := d.Get("project_id").(string)
	key := d.Get("key").(string)
	target := d.Get("target").(string)
	secretName := d.Get("secret_name").(string)

	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf("%s/v4/projects/%s/env/%s?target=%s", "https://api.vercel.com", projectID, key, target),
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

	conf := sdk.NewConfiguration()
	conf.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", token))
	sdkClient := sdk.NewAPIClient(conf)
	_, _, err = sdkClient.SecretsApi.RemoveSecret(ctx, secretName).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	// deliberately ignoring resp.StatusCode here because the Secret is a secondary resource

	d.SetId("")

	return diags
}

func resourceEnvvarDiff(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
	if oldSecretUID, ok := d.GetOk("secret_uid"); ok {
		token := m.(string)

		projectID, _ := d.GetChange("project_id")
		key, _ := d.GetChange("key")
		target, _ := d.GetChange("target")

		conf := sdk.NewConfiguration()
		conf.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", token))
		sdkClient := sdk.NewAPIClient(conf)
		project, _, err := sdkClient.ProjectsApi.GetProjectEnvironmentVariables(ctx, projectID.(string)).Execute()
		if err != nil {
			return err
		}
		for _, envvar := range project.GetEnvs() {
			targetRead := envvar.GetTarget().TargetEnvironment
			if targetRead == nil && envvar.GetTarget().TargetEnvironmentList != nil {
				targetList := *envvar.GetTarget().TargetEnvironmentList
				if len(targetList) > 0 {
					targetRead = &targetList[0]
				}
			}
			if envvar.GetKey() == key && targetRead != nil && strings.EqualFold(string(*targetRead), target.(string)) {
				newSecretUID := envvar.GetValue()
				if oldSecretUID == newSecretUID {
					return nil
				}
				err := d.SetNew("secret_uid", newSecretUID)
				if err != nil {
					return err
				}
				err = d.ForceNew("secret_uid")
				if err != nil {
					return err
				}
				return nil
			}
		}

		return fmt.Errorf("Could not find environment variable in Vercel project")
	}
	return nil
}
