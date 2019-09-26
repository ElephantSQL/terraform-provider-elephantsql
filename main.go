package elephantsql

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
	"github.com/elephantsql/terraform-provider-elephantsql/elephantsql"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return elephantsql.Provider()
		},
	})
}
