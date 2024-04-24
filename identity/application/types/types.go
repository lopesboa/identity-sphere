package types

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
)

type contextKey int

type IdentityManager interface {
	CreateUser(ctx context.Context, user gocloak.User, password string) (*gocloak.User, error)
}

const (
	ContextKeyRequestId contextKey = iota
	ContextKeyClaims    contextKey = iota
)
