package infrastructure

import "github.com/Nerzal/gocloak/v13"

func CreateNewClient(baseUrl string) *gocloak.GoCloak {
	return gocloak.NewClient(baseUrl)
}
