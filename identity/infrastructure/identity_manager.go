package infrastructure

import (
	"context"
	"fmt"

	"github.com/Nerzal/gocloak/v13"
	"github.com/lopesboa/identity-sphere/internal/config"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type identityManager struct {
	BaseUrl             string
	Realm               string
	RestApiClientId     string
	RestApiClientSecret string
	IdOfClient          string
}

type Logger interface {
	Error(v ...interface{})
	Info(v ...interface{})
}

func NewIdentityManager() *identityManager {
	return &identityManager{
		BaseUrl:             viper.GetString("Keycloak.BaseUrl"),
		Realm:               viper.GetString("Keycloak.Realm"),
		RestApiClientId:     viper.GetString("Keycloak.RestApi.ClientId"),
		RestApiClientSecret: viper.GetString("Keycloak.RestApi.ClientSecret"),
		IdOfClient:          viper.GetString("Keycloak.RestApi.IdOfClient"),
	}
}

func (im *identityManager) createNewClient() *gocloak.GoCloak {
	return gocloak.NewClient(im.BaseUrl)
}

func (im *identityManager) loginRestApiClient(ctx context.Context, logger Logger) (*gocloak.JWT, error) {
	client := im.createNewClient()

	token, err := client.LoginClient(ctx, im.RestApiClientId, im.RestApiClientSecret, im.Realm)

	if err != nil {
		logger.Error(err, "unable to login the REST API client")
		return nil, errors.Wrap(err, "unable to login the REST API client")
	}

	return token, nil
}

func (im *identityManager) goClientCreateUser(ctx context.Context, client *gocloak.GoCloak, accessToken string, user gocloak.User, logger Logger) (string, error) {
	userId, err := client.CreateUser(ctx, accessToken, im.Realm, user)

	if err != nil {
		logger.Error(err)
		return "", err
	}

	return userId, nil

}

func (im *identityManager) CreateUser(ctx context.Context, user gocloak.User, password string) (*gocloak.User, error) {
	logger := config.GetLogger("identity manager")

	token, err := im.loginRestApiClient(ctx, logger)

	if err != nil {
		return nil, err
	}

	client := im.createNewClient()

	userId, err := im.goClientCreateUser(ctx, client, token.AccessToken, user, logger)

	if err != nil {
		return nil, err
	}

	err = client.SetPassword(ctx, token.AccessToken, userId, im.Realm, password, false)

	if err != nil {
		return nil, errors.Wrap(err, "unable to set the password  for the user")
	}

	clientRole, err := client.GetClientRole(ctx, token.AccessToken, im.Realm, im.IdOfClient, "cars:read")

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("unable to get role by name: '%v", err))
	}

	err = client.AddClientRolesToUser(ctx, token.AccessToken, im.Realm, im.IdOfClient, userId, []gocloak.Role{*clientRole})

	if err != nil {
		return nil, errors.Wrap(err, "unable to add client role to user")
	}

	newUser, err := client.GetUserByID(ctx, token.AccessToken, im.Realm, userId)

	if err != nil {
		return nil, errors.Wrap(err, "unable to get recently created user")
	}

	return newUser, nil
}

func (im *identityManager) RetrospectToken(ctx context.Context, accessToken string) (*gocloak.IntroSpectTokenResult, error) {
	client := im.createNewClient()

	retrospectToken, err := client.RetrospectToken(ctx, accessToken, im.RestApiClientId, im.RestApiClientSecret, im.Realm)

	if err != nil {
		return nil, errors.Wrap(err, "unable to retrospect token")
	}

	return retrospectToken, nil
}
