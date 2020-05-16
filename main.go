package main

import (
	"github.com/FilippoProjetto/terraform-provider-drp/drp"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: drp.Provider,
	})
}
