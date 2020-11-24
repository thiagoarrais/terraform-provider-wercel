# Project Resource

Represents a [Project](https://vercel.com/docs/api#endpoints/projects) in Vercel infrastructure.

## Example Usage

```hcl
resource "wercel_project" "myproject" {
  name = "myproject"
  repo {
    type        = "gitlab"
    project_url = "https://gitlab.com/myorg/myproject"
  }
  domains = [ "mydomain.com" ]
}
```

## Argument Reference

* `name` - (Required) The project name.
* `repo` - (Required) The repo to import the code from. See [Repositories](#repositories) below.
* `domains` - (Optional) List of domains to associate with this project's deployments.

### Repositories

The `repo` block describes the linked repository where the project is synced from. It
supports the following arguments:

* `type` - (Required) Repo type. Only `gitlab` is supported for now.
* `project_url` - (Required) The main (web) URL for the project/repo.
