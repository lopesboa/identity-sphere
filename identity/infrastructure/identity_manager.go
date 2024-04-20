package infrastructure

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type identityManager struct {
	BaseUrl             string
	Realm               string
	RestApiClientId     string
	RestApiClientSecret string
	client              GoCloakClient
}

func NewIdentityManager(client GoCloakClient) *identityManager {
	return &identityManager{
		BaseUrl:             viper.GetString("Keycloak.BaseUrl"),
		Realm:               viper.GetString("Keycloak.Realm"),
		RestApiClientId:     viper.GetString("Keycloak.ClientId"),
		RestApiClientSecret: viper.GetString("Keycloak.ClientSecret"),
		client:              client,
	}
}

func (im *identityManager) LoginRestApiClient(ctx context.Context) (*gocloak.JWT, error) {
	return im.client.LoginClient(ctx, im.RestApiClientId, im.RestApiClientSecret, im.Realm)
}

func (im *identityManager) CreateUser(ctx context.Context, user gocloak.User, password string, role string) (*gocloak.User, error) {
	token, err := im.LoginRestApiClient(ctx)

	if err != nil {
		return nil, err
	}

	userId, err := im.client.CreateUser(ctx, token.AccessToken, im.Realm, user)

	if err != nil {
		return nil, errors.Wrap(err, "unable to create the user")
	}

	isPassWorkTemporary := false

	err = im.client.SetPassword(ctx, token.AccessToken, userId, im.Realm, password, isPassWorkTemporary)

	if err != nil {
		return nil, errors.Wrap(err, "unable to set the password for the user")
	}

	// TODO: Find a way to asign role to the user based on clientId.

	newUser, err := im.client.GetUserByID(ctx, token.AccessToken, im.Realm, userId)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get recently created user")
	}

	return newUser, nil

}
