package wercel

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func toValidateDiagFunc(validate schema.SchemaValidateFunc) schema.SchemaValidateDiagFunc {
	return func(input interface{}, path cty.Path) diag.Diagnostics {
		var diags diag.Diagnostics
		_, errs := validate(input, "target")
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
				ValidateDiagFunc: toValidateDiagFunc(validation.StringInSlice([]string{"development", "preview", "production"}, true)),
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
	client := &http.Client{}
	token := m.(string)

	var diags diag.Diagnostics

	value := d.Get("value").(string)
	secretName := uuid.New().String()

	reqBody, err := json.Marshal(map[string]string{
		"name":  secretName,
		"value": value,
	})
	if err != nil {
		return diag.FromErr(err)
	}
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/v2/now/secrets", "https://api.vercel.com"),
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return diag.FromErr(err)
	}
	var secret map[string]interface{}
	err = json.Unmarshal(body, &secret)

	secretUID := secret["uid"].(string)
	projectID := d.Get("project_id").(string)
	key := d.Get("key").(string)
	target := d.Get("target").(string)

	reqBody, err = json.Marshal(map[string]string{
		"key":    key,
		"value":  secretUID,
		"target": strings.ToLower(target),
	})
	if err != nil {
		return diag.FromErr(err)
	}
	req, err = http.NewRequest(
		"POST",
		fmt.Sprintf("%s/v4/projects/%s/env", "https://api.vercel.com", projectID),
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

	req, err = http.NewRequest(
		"DELETE",
		fmt.Sprintf("%s/v2/now/secrets/%s", "https://api.vercel.com", secretName),
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err = client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	d.SetId("")

	return diags
}

func resourceEnvvarDiff(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
	if oldSecretUID, ok := d.GetOk("secret_uid"); ok {
		client := &http.Client{}
		token := m.(string)

		projectID, _ := d.GetChange("project_id")
		key, _ := d.GetChange("key")
		target, _ := d.GetChange("target")

		req, err := http.NewRequest(
			"GET",
			fmt.Sprintf("%s/v6/projects/%s/env", "https://api.vercel.com", projectID),
			nil,
		)
		if err != nil {
			return err
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		result := map[string]interface{}{}
		json.NewDecoder(resp.Body).Decode(&result)

		for _, envvarIntf := range result["envs"].([]interface{}) {
			envvar := envvarIntf.(map[string]interface{})
			if envvar["key"].(string) == key && strings.EqualFold(envvar["target"].(string), target.(string)) {
				newSecretUID := envvar["value"].(string)
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
