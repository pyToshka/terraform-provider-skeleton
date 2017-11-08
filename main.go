package main

import (

  "github.com/hashicorp/terraform/plugin"
  "github.com/pyToshka/terraform-provider-skeleton/skeleton"
)

func main() {
  plugin.Serve(&plugin.ServeOpts{
    ProviderFunc: skeleton.Provider})
}