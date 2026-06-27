# Terraform Module Starter

A git repo template for your Terraform module.

Features included here:

- `.terraform-version` [file](https://github.com/tfutils/tfenv#terraform-version-file) with tfenv.
- Testing with [Terratest](https://terratest.gruntwork.io/).
- [Pre-commit hooks with Terraform](https://github.com/antonbabenko/pre-commit-terraform).
- [Makefile for module management with Git submodules](Makefile)

## Overview

This is an opinionated starting point for building a reusable Terraform module.
Fork it (or use it as a GitHub template) and replace the example `in`/`out` logic with your own resources.
It ships with formatting, linting, security scanning, automated tests, and semantic-release wired up out of the box.

## Usage

```hcl
module "example" {
  source = "github.com/anhdle14/terraform-module-template"

  in = "Hello, World!"
}

output "out" {
  value = module.example.out
  # => { lower = "hello, world!", upper = "HELLO, WORLD!" }
}
```

See the [`examples/`](examples) directory for runnable configurations that are exercised by the test suite.

## Documentations

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
| ---- | ------- |
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.0, < 2.0.0 |

## Providers

No providers.

## Modules

No modules.

## Resources

No resources.

## Inputs

| Name | Description | Type | Default | Required |
| ---- | ----------- | ---- | ------- | :------: |
| <a name="input_in"></a> [in](#input\_in) | An example input. | `string` | `"Hello, World!"` | no |

## Outputs

| Name | Description |
| ---- | ----------- |
| <a name="output_out"></a> [out](#output\_out) | An example output. |
<!-- END_TF_DOCS -->

## Contributing

Please refer to the [contribution guidelines](CONTRIBUTING.md) for information on contributing to this module.
