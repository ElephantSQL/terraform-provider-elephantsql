# Terraform provider for ElephantSQL

Setup your ElephantSQL cluster with Terraform

## Install

```sh
git clone https://github.com/elephantsql/terraform-provider.git
cd terraform-provider
make depupdate
make init
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
  value = "${elephantsql_instance.sql_hippo.url}"
}
```



