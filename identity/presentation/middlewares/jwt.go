package middlewares

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"

	"github.com/Nerzal/gocloak/v13"
	contribJwt "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	golangJwt "github.com/golang-jwt/jwt/v5"
	"github.com/lopesboa/identity-sphere/identity/application/types"
	"github.com/spf13/viper"
)

type TokenRetrospect interface {
	RetrospectToken(ctx context.Context, accessToken string) (*gocloak.IntroSpectTokenResult, error)
}

func NewJwtMiddleware(tokenRetrospect TokenRetrospect) fiber.Handler {
	base64Str := viper.GetString("Keycloak.RealmRS256PublicKey")
	publicKey, err := parseKeycloakRSAPublicKey(base64Str)

	if err != nil {
		panic(err)
	}

	return contribJwt.New(contribJwt.Config{
		SigningKey: contribJwt.SigningKey{
			JWTAlg: contribJwt.RS256,
			Key:    publicKey,
		},
		SuccessHandler: func(ctx *fiber.Ctx) error {
			return successHandler(ctx, tokenRetrospect)
		},
	})
}

func successHandler(fiberCtx *fiber.Ctx, tokenRetrospect TokenRetrospect) error {
	jwtToken := fiberCtx.Locals("user").(*golangJwt.Token)
	claims := jwtToken.Claims.(golangJwt.MapClaims)

	var ctx = fiberCtx.UserContext()
	var contextWithClaims = context.WithValue(ctx, types.ContextKeyClaims, claims)
	fiberCtx.SetUserContext(contextWithClaims)

	retrospectTokenResult, err := tokenRetrospect.RetrospectToken(ctx, jwtToken.Raw)

	if err != nil {
		panic(err)
	}

	if !*retrospectTokenResult.Active {
		return fiberCtx.Status(fiber.StatusUnauthorized).SendString("invalid credentials")
	}

	return fiberCtx.Next()
}

func parseKeycloakRSAPublicKey(base64String string) (*rsa.PublicKey, error) {
	buf, err := base64.StdEncoding.DecodeString(base64String)

	if err != nil {
		return nil, err
	}

	parsedKey, err := x509.ParsePKIXPublicKey(buf)

	if err != nil {
		return nil, err
	}

	publicKey, ok := parsedKey.(*rsa.PublicKey)

	if ok {
		return publicKey, nil
	}

	return nil, fmt.Errorf("unexpected key type %T", publicKey)
}
