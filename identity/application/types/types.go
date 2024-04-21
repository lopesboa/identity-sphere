package types

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
)

type IdentityManager interface {
	CreateUser(ctx context.Context, user gocloak.User, password string) (*gocloak.User, error)
}
