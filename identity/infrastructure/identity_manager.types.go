package infrastructure

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
)

type GoCloakClient interface {
	LoginClient(ctx context.Context, clientID string, clientSecret string, realm string, scopes ...string) (*gocloak.JWT, error)
	CreateUser(ctx context.Context, token string, realm string, user gocloak.User) (string, error)
	SetPassword(ctx context.Context, token string, userID string, realm string, password string, temporary bool) error
	GetUserByID(ctx context.Context, accessToken string, realm string, userID string) (*gocloak.User, error)
	GetClientRole(ctx context.Context, token string, realm string, idOfClient string, roleName string) (*gocloak.Role, error)
}
