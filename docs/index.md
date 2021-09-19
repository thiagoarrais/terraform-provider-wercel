# Wercel Provider

The Wercel Provider is an unofficial Terraform Provider for [Vercel](https://vercel.com).
It can be used to create and manage Vercel projects.

Tokens can be obtained from [Vercel's settings page](https://vercel.com/account/tokens).

## Deprecation notice

This is discontinued and no longer maintained. A good provider for Vercel is [chronark's](https://registry.terraform.io/providers/chronark/vercel/).
You may want to take a look at it.

## Example Usage

```hcl
provider "wercel" {
    token = "yourVercelToken"
}
```

## Argument Reference

* `token`  - (Required) a valid Vercel token that shall be used to manage infrastructure.
