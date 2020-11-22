package wercel

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/thiagoarrais/terraform-provider-wercel/sdk"
)

func TestAccWercelProject_basic(t *testing.T) {
	var project sdk.Project
	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWercelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccWercelProject(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWercelProjectExists("wercel_project.myproject", &project),
				),
			},
		},
	})
}

func testAccWercelProject(name string) string {
	return fmt.Sprintf(`
provider "wercel" {
  token = "%s"
}

resource "wercel_project" "myproject" {
  name = "%s"
  repo {
  	type        = "gitlab"
  	project_url = "https://gitlab.com/arrais-tfvercel/hello-world"
  }
}`, os.Getenv("VERCEL_TOKEN"), name)
}

// testAccCheckWercelProjectExists uses the Vercel API directly to retrieve
// the Project data, and stores it in the provided
// *sdk.Project
func testAccCheckWercelProjectExists(resourceName string, project *sdk.Project) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// retrieve the resource by name from state
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Project ID is not set")
		}

		conf := sdk.NewConfiguration()
		conf.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("VERCEL_TOKEN")))
		sdkClient := sdk.NewAPIClient(conf)

		result, _, err := sdkClient.ProjectsApi.GetProjectById(context.Background(), rs.Primary.ID).Execute()
		if err != nil {
			return errorFromSDKErr(err)

		}

		*project = result

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
