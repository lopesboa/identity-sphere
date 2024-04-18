package identity

import (
	identity "github.com/lopesboa/identity-sphere/identity/domain/repositories"
	"github.com/spf13/viper"
)

func NewIdentityManager() *identity.IdentityManager {
	return &identity.IdentityManager{
		BaseUrl:             viper.GetString("Keycloak.BaseUrl"),
		Realm:               viper.GetString("Keycloak.Realm"),
		RestApiClientId:     viper.GetString("Keycloak.ClientId"),
		RestApiClientSecret: viper.GetString("Keycloak.ClientSecret"),
	}
}
