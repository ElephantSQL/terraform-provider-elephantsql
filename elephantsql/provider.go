package main

import (
	"github.com/84codes/go-api/api"
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"apikey": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ELEPHANTSQL_APIKEY", nil),
				Description: "Key used to authentication to the ElephantSQL API",
			},
			"baseurl": &schema.Schema{
				Type:        schema.TypeString,
				Default:     "https://customer.elephantsql.com",
				Optional:    true,
				Description: "Base URL to ElephantSQL website",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"elephantsql_instance": resourceInstance(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return api.New(d.Get("baseurl").(string), d.Get("apikey").(string)), nil
}
