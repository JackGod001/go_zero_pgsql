package casdoor

import "github.com/casdoor/casdoor-go-sdk/casdoorsdk"

var CasdoorEndpoint = "http://127.0.0.1:8000"
var ClientId = "4447b4325613599531b3"
var ClientSecret = "dd8982f7046ccba1bbd7851d5c1ece4e52bf039d"
var CasdoorOrganization = "built-in"
var CasdoorApplication = "next_base_auth"

var JwtPublicKey string

func init() {
	casdoorsdk.InitConfig(CasdoorEndpoint, ClientId, ClientSecret, JwtPublicKey, CasdoorOrganization, CasdoorApplication)
}
