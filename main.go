package main

import (
	accessmanager "github.com/jralmaraz/forgerock-am-terraform-provider"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return accessmanager.Provider()
		},
	})
}
