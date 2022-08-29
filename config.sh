#!/usr/bin/env sh
set -e
go fmt ./...
go mod tidy
make install
rm -rf examples/.terraform examples/.terraform.lock.hcl examples/terraform.tfstate
cd examples
terraform init
terraform apply