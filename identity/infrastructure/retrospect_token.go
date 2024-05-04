package infrastructure

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
	"github.com/pkg/errors"
)

func RetrospectToken(ctx context.Context, accessToken string) (*gocloak.IntroSpectTokenResult, error) {
	im := NewIdentityManager()

	client := im.createNewClient()

	retrospectToken, err := client.RetrospectToken(ctx, accessToken, im.RestApiClientId, im.RestApiClientSecret, im.Realm)

	if err != nil {
		return nil, errors.Wrap(err, "unable to retrospect token")
	}

	return retrospectToken, nil
}
