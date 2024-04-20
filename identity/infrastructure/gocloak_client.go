package infrastructure

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
)

type gocloakClient struct {
	client *gocloak.GoCloak
}

func NewGoCloakClient(baseUrl string, options ...func(*gocloak.GoCloak)) GoCloakClient {
	return &gocloakClient{
		client: gocloak.NewClient(baseUrl),
	}
}

func (gc *gocloakClient) CreateUser(ctx context.Context, token string, realm string, user gocloak.User) (string, error) {
	return gc.client.CreateUser(ctx, token, realm, user)
}

func (gc *gocloakClient) LoginClient(ctx context.Context, clientID string, clientSecret string, realm string, scopes ...string) (*gocloak.JWT, error) {
	return gc.LoginClient(ctx, clientID, clientSecret, realm)
}

func (gc *gocloakClient) GetUserByID(ctx context.Context, accessToken string, realm string, userID string) (*gocloak.User, error) {
	return gc.client.GetUserByID(ctx, accessToken, realm, userID)
}

func (gc *gocloakClient) SetPassword(ctx context.Context, token string, userID string, realm string, password string, temporary bool) error {
	return gc.client.SetPassword(ctx, token, userID, realm, password, temporary)
}

func (gc *gocloakClient) GetClientRole(ctx context.Context, token string, realm string, idOfClient string, roleName string) (*gocloak.Role, error) {
	return gc.client.GetClientRole(ctx, token, realm, idOfClient, roleName)
}
