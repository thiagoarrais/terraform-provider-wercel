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
				Config: testAccWercelGitlabProject(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWercelProjectExists("wercel_project.myproject", &project),
					testAccCheckEquals("Unexpected project name", func() interface{} { return project.GetName() }, rName),
					testAccCheckEquals("Unexpected alias size", func() interface{} { return len(project.GetAlias()) }, 2),
					testAccCheckEquals("Unexpected alias", func() interface{} { return project.GetAlias()[1].GetDomain() }, rName+"-extradomain.vercel.app"),
					testAccCheckEquals("Unexpected link type", func() interface{} { return project.Link.GetType() }, "gitlab"),
					testAccCheckEquals("Unexpected link project URL", func() interface{} { return project.Link.GetProjectUrl() }, "https://gitlab.com/arrais-tfvercel/hello-world"),
					resource.TestCheckResourceAttr("wercel_project.myproject", "name", rName),
					resource.TestCheckResourceAttr("wercel_project.myproject", "repo.0.type", "gitlab"),
					resource.TestCheckResourceAttr("wercel_project.myproject", "repo.0.project_url", "https://gitlab.com/arrais-tfvercel/hello-world"),
					resource.TestCheckResourceAttr("wercel_project.myproject", "domains.#", "1"),
					resource.TestCheckResourceAttr("wercel_project.myproject", "domains.0", rName+"-extradomain.vercel.app"),
				),
			},
			{
				Config: testAccWercelGitHubProject_withoutDomains(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWercelProjectExists("wercel_project.myproject", &project),
					testAccCheckEquals("Unexpected alias size", func() interface{} { return len(project.GetAlias()) }, 1),
					resource.TestCheckResourceAttr("wercel_project.myproject", "repo.0.type", "github"),
					resource.TestCheckResourceAttr("wercel_project.myproject", "repo.0.project_url", "https://github.com/thiagoarrais/tfvercel"),
					resource.TestCheckResourceAttr("wercel_project.myproject", "domains.#", "0"),
				),
			},
		},
	})
}

func TestAccWercelProject_fromGitHubToGitlab(t *testing.T) {
	var project sdk.Project
	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWercelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccWercelGitHubProject(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWercelProjectExists("wercel_project.myproject", &project),
					testAccCheckEquals("Unexpected link type", func() interface{} { return project.Link.GetType() }, "github"),
					testAccCheckEquals("Unexpected link org", func() interface{} { return project.Link.GetOrg() }, "thiagoarrais"),
					testAccCheckEquals("Unexpected link repo", func() interface{} { return project.Link.GetRepo() }, "tfvercel"),
					resource.TestCheckResourceAttr("wercel_project.myproject", "name", rName),
					resource.TestCheckResourceAttr("wercel_project.myproject", "repo.0.type", "github"),
					resource.TestCheckResourceAttr("wercel_project.myproject", "repo.0.project_url", "https://github.com/thiagoarrais/tfvercel"),
				),
			},
			{
				Config: testAccWercelGitlabProject(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWercelProjectExists("wercel_project.myproject", &project),
					resource.TestCheckResourceAttr("wercel_project.myproject", "repo.0.type", "gitlab"),
					resource.TestCheckResourceAttr("wercel_project.myproject", "repo.0.project_url", "https://gitlab.com/arrais-tfvercel/hello-world"),
				),
			},
		},
	})
}

func TestAccWercelProject_withRootDirectory(t *testing.T) {
	var project sdk.Project
	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckWercelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccWercelGitHubProjectDirectoryA(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWercelProjectExists("wercel_project.myproject", &project),
					testAccCheckEquals("Unexpected link type", func() interface{} { return project.Link.GetType() }, "github"),
					testAccCheckEquals("Unexpected link org", func() interface{} { return project.Link.GetOrg() }, "thiagoarrais"),
					testAccCheckEquals("Unexpected link repo", func() interface{} { return project.Link.GetRepo() }, "tfvercel"),
					testAccCheckEquals("Unexpected root directory", func() interface{} { return project.HasRootDirectory() }, true),
					testAccCheckEquals("Unexpected root directory", func() interface{} { return project.GetRootDirectory() }, "experiment-a"),
					resource.TestCheckResourceAttr("wercel_project.myproject", "name", rName),
					resource.TestCheckResourceAttr("wercel_project.myproject", "repo.0.type", "github"),
					resource.TestCheckResourceAttr("wercel_project.myproject", "repo.0.project_url", "https://github.com/thiagoarrais/tfvercel"),
					resource.TestCheckResourceAttr("wercel_project.myproject", "root", "experiment-a"),
				),
			},
			{
				Config: testAccWercelGitHubProjectDirectoryB(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWercelProjectExists("wercel_project.myproject", &project),
					testAccCheckEquals("Unexpected root directory", func() interface{} { return project.HasRootDirectory() }, true),
					testAccCheckEquals("Unexpected root directory", func() interface{} { return project.GetRootDirectory() }, "experiment-b"),
					resource.TestCheckResourceAttr("wercel_project.myproject", "root", "experiment-b"),
				),
			},
			{
				Config: testAccWercelGitHubProject(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWercelProjectExists("wercel_project.myproject", &project),
					testAccCheckEquals("Unexpected root directory", func() interface{} { return project.HasRootDirectory() }, false),
					resource.TestCheckResourceAttr("wercel_project.myproject", "root", ""),
				),
			},
		},
	})
}

func testAccWercelGitlabProject(name string) string {
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
  domains = [ "%s-extradomain.vercel.app" ]
}`, os.Getenv("VERCEL_TOKEN"), name, name)
}

func testAccWercelGitHubProject(name string) string {
	return fmt.Sprintf(`
provider "wercel" {
  token = "%s"
}

resource "wercel_project" "myproject" {
  name = "%s"
  repo {
    type        = "github"
    project_url = "https://github.com/thiagoarrais/tfvercel"
  }
  domains = [ "%s-extradomain.vercel.app" ]
}`, os.Getenv("VERCEL_TOKEN"), name, name)
}

func testAccWercelGitHubProject_withoutDomains(name string) string {
	return fmt.Sprintf(`
provider "wercel" {
  token = "%s"
}

resource "wercel_project" "myproject" {
  name = "%s"
  repo {
    type        = "github"
    project_url = "https://github.com/thiagoarrais/tfvercel"
  }
}`, os.Getenv("VERCEL_TOKEN"), name)
}

func testAccWercelGitHubProjectDirectoryA(name string) string {
	return fmt.Sprintf(`
provider "wercel" {
  token = "%s"
}

resource "wercel_project" "myproject" {
  name = "%s"
  repo {
    type        = "github"
    project_url = "https://github.com/thiagoarrais/tfvercel"
  }
  root = "experiment-a"
}`, os.Getenv("VERCEL_TOKEN"), name)
}

func testAccWercelGitHubProjectDirectoryB(name string) string {
	return fmt.Sprintf(`
provider "wercel" {
  token = "%s"
}

resource "wercel_project" "myproject" {
  name = "%s"
  repo {
    type        = "github"
    project_url = "https://github.com/thiagoarrais/tfvercel"
  }
  root = "experiment-b"
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
