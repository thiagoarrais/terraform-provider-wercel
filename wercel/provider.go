package wercel

import (
	"context"
	"fmt"

	"github.com/thiagoarrais/terraform-provider-wercel/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"wercel_project":              resourceProject(),
			"wercel_environment_variable": resourceEnvvar(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

type config struct {
	sdkClient *sdk.APIClient
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	token := d.Get("token").(string)

	conf := sdk.NewConfiguration()
	conf.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", token))
	cfg := config{
		sdkClient: sdk.NewAPIClient(conf),
	}

	return cfg, diags
}
