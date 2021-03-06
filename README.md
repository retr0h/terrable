# Terrable

[![Go](https://github.com/retr0h/terraform-provider-terrable/actions/workflows/go.yml/badge.svg)](https://github.com/retr0h/terraform-provider-terrable/actions/workflows/go.yml)
[![Docker Image CI](https://github.com/retr0h/terraform-provider-terrable/actions/workflows/docker-image.yml/badge.svg)](https://github.com/retr0h/terraform-provider-terrable/actions/workflows/docker-image.yml)
[![Linter](https://github.com/retr0h/terraform-provider-terrable/actions/workflows/linter.yml/badge.svg)](https://github.com/retr0h/terraform-provider-terrable/actions/workflows/linter.yml)
[![codecov](https://codecov.io/gh/retr0h/terraform-provider-terrable/branch/main/graph/badge.svg?token=69E6QZNRVX)](https://codecov.io/gh/retr0h/terraform-provider-terrable)

Terrable ~ Terraform ~ Ansible

This project is a [terraform](http://www.terraform.io/) provider for
on-host configuration management powered by Terraform.

## Rationale

Terraform already provides a [configuration language][] and state management.
This project is an experiment in on-host configuration.

[Ansible is slow][] for our use case, even when running locally or using an
alternative [strategy plugin][].  This project is not intended as a replacement
to Ansible, or as a mechanism to [converge][] remote hosts, rather an experiment
with our specific use cases.

[configuration language]: https://github.com/hashicorp/hcl
[Ansible is slow]: https://github.com/ansible/ansible/pull/72184
[strategy plugin]: https://mitogen.networkgenomics.com/ansible_detailed.html
[converge]: https://verticalsysadmin.com/blog/idempotence-vs-convergence-in-configuration-management/

## Requirements

* Terraform 0.13.x
* Go 1.15

## Usage

[Resources](docs/resources.md)

### Screenshot

![Usage](img/demo.png?raw=true "Usage")

## Testing

Run unit tests

```sh
make test
```

Run integration tests

```sh
make build-docker test-integration
```

## Developing

### Dependencies for building from source

If you need to build from source, you should have a working Go environment setup.
If not check out the Go [getting started](http://golang.org/doc/install) guide.

This project uses [Go Modules](https://github.com/golang/go/wiki/Modules) for dependency management.
To fetch all dependencies run `make mod` inside this repository.

### Build

```sh
make build
```

The binary will then be available at `build/$(GOOS)_$(GOARCH)/$(PLUGIN_NAME)_v$(VERSION)`

### Install

```sh
make install
```

This will place the binary under `$(HOME)/.terraform.d/plugins/$(HOSTNAME)/$(USER)/$(NAME)/$(VERSION)/$(GOOS)_$(GOARCH)`.
After installing you will need to run `terraform init` in any project using the plugin.

## License

The [MIT] License.

[MIT]: LICENSE
