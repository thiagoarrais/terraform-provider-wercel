package wercel

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEnvvar() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		CreateContext: resourceEnvvarCreate,
		ReadContext:   resourceEnvvarRead,
		DeleteContext: resourceEnvvarDelete,
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
		},
	}
}

func resourceEnvvarCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{}
	token := m.(string)

	var diags diag.Diagnostics

	value := d.Get("value").(string)

	reqBody, err := json.Marshal(map[string]string{
		"name":  uuid.New().String(),
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

	reqBody, err = json.Marshal(map[string]string{
		"key":   key,
		"value": secretUID,
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

	return diags
}

func resourceEnvvarRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceEnvvarDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}
