package wercel

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/thiagoarrais/terraform-provider-wercel/sdk"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"wercel": testAccProvider,
	}
}

func testAccCheckWercelDestroy(s *terraform.State) error {
	conf := sdk.NewConfiguration()
	conf.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("VERCEL_TOKEN")))
	sdkClient := sdk.NewAPIClient(conf)

	projectsResult, _, err := sdkClient.ProjectsApi.GetProjects(context.Background()).Execute()
	if err != nil {
		return fmt.Errorf("Could not check Vercel. There may be dangling resources. Reason: %s", errorFromSDKErr(err).Error())
	}

	numProjects := len(projectsResult.GetProjects())
	if numProjects > 0 {
		return fmt.Errorf("%d projects found. Resource destruction did not work.", numProjects)
	}

	secretsResult, _, err := sdkClient.SecretsApi.ListSecrets(context.Background()).Execute()
	if err != nil {
		return fmt.Errorf("Could not check Vercel. There may be dangling resources. Reason: %s", errorFromSDKErr(err).Error())
	}

	numSecrets := len(secretsResult.GetSecrets())
	if numSecrets > 0 {
		return fmt.Errorf("%d secrets found. Resource destruction did not work.", numSecrets)
	}

	return nil
}

func testAccCheckEquals(msg string, actualOp func() interface{}, expected interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		actual := actualOp()
		if actual != expected {
			return fmt.Errorf("%s. Expected: %#v. Actual: %#v", msg, expected, actual)

		}
		return nil
	}
}

// testAccPreCheck validates the necessary test API keys exist
// in the testing environment
func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("VERCEL_TOKEN"); v == "" {
		t.Fatal("VERCEL_TOKEN must be set for acceptance tests")
	}
}
