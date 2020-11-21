package wercel

import (
	"context"
	"fmt"
	"os"

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

	return nil
}
