package types

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
)

type contextKey int

type LoginResponse struct {
	AccessToken  string
	ExpiresIn    int
	RefreshToken string
	Scope        string
}

type IdentityManager interface {
	CreateUser(ctx context.Context, user gocloak.User) error
	// Login(ctx context.Context, ) (*LoginResponse, error)
}

const (
	ContextKeyRequestId contextKey = iota
	ContextKeyClaims    contextKey = iota
)
