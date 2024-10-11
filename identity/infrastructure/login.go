package infrastructure

import (
	"context"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type LoginParams struct {
	Username string
	Password string
}

type LoginResponse struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int
	Scope        string
}

type loginManagerConfig struct {
	BaseUrl             string
	RestApiClientId     string
	RestApiClientSecret string
	Realm               string
}

type LoginClient interface {
	Login(ctx context.Context, params LoginParams) (*LoginResponse, error)
}

func NewLoginManager() *loginManagerConfig {
	return &loginManagerConfig{
		BaseUrl:             viper.GetString("Keycloak.BaseUrl"),
		Realm:               viper.GetString("Keycloak.Realm"),
		RestApiClientId:     viper.GetString("Keycloak.RestApi.ClientId"),
		RestApiClientSecret: viper.GetString("Keycloak.RestApi.ClientSecret"),
	}
}

func (l *LoginResponse) GetAccessToken() string {
	return l.AccessToken
}

func (l *LoginResponse) GetRefreshToken() string {
	return l.RefreshToken
}

func (lm *loginManagerConfig) Login(ctx context.Context, params LoginParams) (*LoginResponse, error) {
	gc := CreateNewClient(lm.BaseUrl)

	token, err := gc.Login(ctx, lm.RestApiClientId, lm.RestApiClientSecret, lm.Realm, params.Username, params.Password)

	if err != nil {
		return nil, errors.Wrap(err, "unable to login the REST API client")
	}

	return &LoginResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresIn:    token.ExpiresIn,
		Scope:        token.Scope,
	}, nil

}
