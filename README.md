# Terraform provider for ElephantSQL

Setup your ElephantSQL cluster with Terraform. Create your account at https://customer.elephantsql.com/ to get an API key.

## Example

```hcl
terraform {
  required_providers {
    elephantsql = {
      source = "elephantsql/elephantsql"
    }
  }
}

resource "elephantsql_instance" "sql_hippo" {
  name   = "terraform-provider-test"
  plan   = "hippo"
  region = "amazon-web-services::us-east-1"
}

output "psql_url" {
  sensitive = true
  value = elephantsql_instance.sql_hippo.url
}
```

```bash
export ELEPHANTSQL_APIKEY=<your-key>

terraform init

# Create the database instance
terraform apply -auto-approve

# View the database URL
terraform output -raw psql_url
```

## Use from source

```sh
git clone https://github.com/ElephantSQL/terraform-provider-elephantsql.git
cd terraform-provider-elephantsql
make local-install
```

You can use [`dev_overrides`](https://developer.hashicorp.com/terraform/cli/config/config-file#development-overrides-for-provider-developers) to make Terraform use your local build:

```hcl
# save as $HOME/elephantsql-provider-dev.tfrc
provider_installation {
  dev_overrides {
    "elephantsql/elephantsql" = "/home/foo/terraform-provider-elephantsql"
  }
}
```

```sh
TERRAFORM_CONFIG=$HOME/elephantsql-provider-dev.tfrc terraform apply -auto-approve
```
