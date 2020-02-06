package main

import (
	"github.com/elephantsql/terraform-provider-elephantsql/elephantsql"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return elephantsql.Provider()
		},
	})
}
