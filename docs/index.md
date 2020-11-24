# Wercel Provider

The Wercel Provider is an unofficial Terraform Provider for [Vercel](https://vercel.com).
It can be used to create and manage Vercel projects.

Tokens can be obtained from [Vercel's settings page](https://vercel.com/account/tokens).

## Example Usage

```hcl
provider "wercel" {
    token = "yourVercelToken"
}
```

## Argument Reference

* `token`  - (Required) a valid Vercel token that shall be used to manage infrastructure.
