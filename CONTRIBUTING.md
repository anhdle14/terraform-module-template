# CONTRIBUTING

This document provides guidelines for contributing to the module.

## Prerequisites

- [tfenv](https://github.com/tfutils/tfenv)
- [terraform](https://learn.hashicorp.com/terraform/getting-started/install#installing-terraform)
  - Can be skipped when `tfenv` is installed.
- [terraform-docs](https://github.com/terraform-docs/terraform-docs)
- [terratest](https://terratest.gruntwork.io)
- [golang](https://golang.org/doc/install#install)
- [staticcheck](https://staticcheck.io)
- [pre-commit](https://pre-commit.com/#install)
- [git-cz](https://www.npmjs.com/package/git-cz)
- [infracost](https://www.infracost.io)
- [checkov](https://www.checkov.io)
- [trivy](https://trivy.dev) (replaces the now-deprecated tfsec)
- [makefile](https://www.gnu.org/software/make/manual/make.html)

### Tests

- Test dependencies are declared in `tests/go.mod` and fetched automatically by Go modules.
- Tests live in the `tests/` directory and run against the configurations in `examples/`.
- Run the whole suite from the repo root:

  ```sh
  make test
  ```


## Manage Terraform modules with Git submodules

- **Add**: Will use `git submodule add`

  ```
  make add module=$NAME module_url=$GIT_URL
  ```


- **Update**: Will use `git checkout` on submodule

  ```
  make update module=$NAME version=$VERSION
  ```

- **Delete**: Will use `git submodule deinit`

  ```
  make delete module=$NAME
  ```

- **Delete all**: Will use `git submodule deinit` but for all modules

  ```
  make delete-all
  ```

## References

- [Testing Hashicorp Terraform](https://www.hashicorp.com/blog/testing-hashicorp-terraform)
- [Awesome Terraform](https://github.com/shuaibiyy/awesome-terraform)
