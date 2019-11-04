# Terraform provider for ElephantSQL

Setup your ElephantSQL cluster with Terraform

## Install

```sh
git clone https://github.com/ElephantSQL/terraform-provider-elephantsql.git
cd terraform-provider-elephantsql
make depupdate
make install
```

Now the provider is installed in the terraform plugins folder and ready to be used.

## Example

```hcl
provider "elephantsql" {}

resource "elephantsql_instance" "sql_hippo" {
  name   = "terraform-provider-test"
  plan   = "hippo"
  region = "amazon-web-services::us-east-1"
}

output "psql_url" {
  value = elephantsql_instance.sql_hippo.url
}
```

## Working with HashiCorp Enterprise remote state

To work with HashiCorp Enterprise remote state you need to build `linux_amd64` binary and unpack it.

```sh
make release
cd bin/release/linux/amd64
gunzip -c /terraform-provider-elephantsql_linux_amd64.tar.gz | tar xopf -
```

You need to move it to `terraform.d/plugins/linux_64` directory (relational, inside your Terraform directory).
```sh
mv terraform-provider-elephantsql your_working_directory/terraform.d/plugins/linux_64
```
Your cloud should be now working properly.

Keep in mind it's dynamically changing, so please follow those [instructions](https://www.terraform.io/docs/cloud/run/index.html#custom-and-community-providers).
