provider "elephantsql" {}

resource "elephantsql_instance" "sql_hippo" {
  name   = "terraform-provider-test"
  plan   = "hippo"
  region = "amazon-web-services::us-east-1"
}

output "psql_url" {
  value = "${elephantsql_instance.sql_hippo.url}"
}
