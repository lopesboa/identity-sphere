package infrastructure

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
)

type GoCloakClient interface {
	loginClient(ctx context.Context, clientID string, clientSecret string, realm string, scopes ...string) (*gocloak.JWT, error)
	createUser(ctx context.Context, token string, realm string, user gocloak.User) (string, error)
	setPassword(ctx context.Context, token string, userID string, realm string, password string, temporary bool) error
	getUserByID(ctx context.Context, accessToken string, realm string, userID string) (*gocloak.User, error)
	getClientRole(ctx context.Context, token string, realm string, idOfClient string, roleName string) (*gocloak.Role, error)
}

type gocloakClient struct {
	client *gocloak.GoCloak
}
