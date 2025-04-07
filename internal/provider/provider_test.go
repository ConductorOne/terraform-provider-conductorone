package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const (
	providerConfig = `
provider "conductorone" {
  server_url    = ""
  client_id     = ""
  client_secret = ""
}
`
)

var (
	testAccProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"conductorone": providerserver.NewProtocol6WithError(New("test")()),
	}
)
