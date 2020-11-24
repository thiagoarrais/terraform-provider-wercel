# Environament Variable Resource

Represents a single environment variable in a Vercel [Project](https://vercel.com/docs/api#endpoints/projects).

The Wercel Terraform Provider only supports variables of type Secret and in one environment.

## Example Usage

```hcl
resource "wercel_environment_variable" "myvar" {
  project_id    = wercel_project.myproject.id
  key           = "MYVAR"
  value         = "the value"
  target        = "production"
}
```

## Argument Reference

* `project_id` - (Required) The Vercel-given project ID. Can be obtained from a [Project][project] Resource's `id`
* `key` - (Required) The variable key.
* `value` - (Required) The variable value.
* `target` - (Optional) The environment where the variable should be made available. Possible values: `development`, `preview` or `production`. Default: `production`

## Attribute Reference

* `secret_name` - The name of the associated secret
* `secret_uid` - The UID of the associated secret