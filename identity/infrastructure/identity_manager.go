package infrastructure

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
	"github.com/lopesboa/identity-sphere/internal/tools"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type identityManager struct {
	BaseUrl             string
	Realm               string
	RestApiClientId     string
	RestApiClientSecret string
	IdOfClient          string
	RedirectURI         string
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
		RedirectURI:         viper.GetString("Keycloak.RestApi.RedirectURI"),
	}
}

func (im *identityManager) loginRestApiClient(ctx context.Context, logger Logger) (*gocloak.JWT, error) {
	client := CreateNewClient(im.BaseUrl)

	token, err := client.LoginClient(ctx, im.RestApiClientId, im.RestApiClientSecret, im.Realm)

	if err != nil {
		logger.Error(err, "unable to login the REST API client")
		return nil, errors.Wrap(err, "unable to login the REST API client")
	}

	return token, nil
}

func (im *identityManager) CreateUser(ctx context.Context, user gocloak.User) error {
	logger := tools.GetLogger("identity manager")

	token, err := im.loginRestApiClient(ctx, logger)

	if err != nil {
		return err
	}

	client := CreateNewClient(im.BaseUrl)

	userId, err := client.CreateUser(ctx, token.AccessToken, im.Realm, user)

	if err != nil {
		logger.Error(err, userId)
		return err
	}

	return nil
}
