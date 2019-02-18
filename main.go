package main

import (
	"./oktaccs"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: oktaccs.Provider})
}
