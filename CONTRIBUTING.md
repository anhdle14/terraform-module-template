# CONTRIBUTING

This document provides guidelines for contributing to the module.

## Prerequisites

- [tfenv](https://github.com/tfutils/tfenv)
- [terraform](https://learn.hashicorp.com/terraform/getting-started/install#installing-terraform)
  - Can be skipped when `tfenv` is installed.
- [terraform-docs](https://github.com/segmentio/terraform-docs)
- [terratest](https://terratest.gruntwork.io)
- [golang](https://golang.org/doc/install#install)
- [staticcheck](https://staticcheck.io)
- [pre-commit](https://pre-commit.com/#install)
- [standard-version](https://github.com/conventional-changelog/standard-version)
- [git-cz](https://www.npmjs.com/package/git-cz)
- [infracost](https://www.infracost.io)
- [checkov](https://www.checkov.io)
- [tfsec](https://tfsec.dev)
- [makefile](https://www.gnu.org/software/make/manual/make.html)

### Tests

- Configure golang deps for tests

  ```sh
  go get github.com/gruntwork-io/terratest/modules/terraform
  go get github.com/stretchr/testify/assert
  ```

- Tests are available in `test` directory
- In the test directory, run the below command

  ```sh
  go test
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

- ***Delete**: Will use `git submodule deinit`

  ```
  make update module=$NAME
  ```

- **Clean**: Will use `git submodule deinit` but for all modules

  ```
  make clean
  ```

## References

- [Testing Hashicorp Terraform](https://www.hashicorp.com/blog/testing-hashicorp-terraform)
- [Awesome Terraform](https://github.com/shuaibiyy/awesome-terraform)
