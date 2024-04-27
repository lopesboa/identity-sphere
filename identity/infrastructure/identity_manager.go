package infrastructure

import (
	"context"
	"fmt"

	"github.com/Nerzal/gocloak/v13"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type identityManager struct {
	BaseUrl             string
	Realm               string
	RestApiClientId     string
	RestApiClientSecret string
}

func NewIdentityManager() *identityManager {
	return &identityManager{
		BaseUrl:             viper.GetString("Keycloak.BaseUrl"),
		Realm:               viper.GetString("Keycloak.Realm"),
		RestApiClientId:     viper.GetString("Keycloak.RestApi.ClientId"),
		RestApiClientSecret: viper.GetString("Keycloak.RestApi.ClientSecret"),
	}
}

func (im *identityManager) loginRestApiClient(ctx context.Context) (*gocloak.JWT, error) {
	client := gocloak.NewClient(im.BaseUrl)

	token, err := client.LoginClient(ctx, im.RestApiClientId, im.RestApiClientSecret, im.Realm)

	if err != nil {
		return nil, errors.Wrap(err, "unable to login the REST API client")
	}

	return token, nil
}

func (im *identityManager) CreateUser(ctx context.Context, user gocloak.User, password string) (*gocloak.User, error) {
	token, err := im.loginRestApiClient(ctx)

	if err != nil {
		return nil, err
	}

	client := gocloak.NewClient(im.BaseUrl)

	userId, err := client.CreateUser(ctx, token.AccessToken, im.Realm, user)

	if err != nil {
		return nil, errors.Wrap(err, "unable to create the user")
	}

	err = client.SetPassword(ctx, token.AccessToken, userId, im.Realm, password, false)

	if err != nil {
		return nil, errors.Wrap(err, "unable to set the password  for the user")
	}

	clientRole, err := client.GetClientRole(ctx, token.AccessToken, im.Realm, im.RestApiClientId, "cars:read")

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("unable to get role by name: '%v"))
	}

	err = client.AddClientRolesToUser(ctx, token.AccessToken, im.Realm, im.RestApiClientId, userId, []gocloak.Role{*clientRole})

	if err != nil {
		return nil, errors.Wrap(err, "unable to add client role to user")
	}

	newUser, err := client.GetUserByID(ctx, token.AccessToken, im.Realm, userId)

	if err != nil {
		return nil, errors.Wrap(err, "unable to get recently created user")
	}

	return newUser, nil
}
