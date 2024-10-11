package infrastructure

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
	"github.com/pkg/errors"
)

func (im *identityManager) RetrospectToken(ctx context.Context, accessToken string) (*gocloak.IntroSpectTokenResult, error) {

	client := CreateNewClient(im.BaseUrl)

	retrospectToken, err := client.RetrospectToken(ctx, accessToken, im.RestApiClientId, im.RestApiClientSecret, im.Realm)

	if err != nil {
		return nil, errors.Wrap(err, "unable to retrospect token")
	}

	return retrospectToken, nil
}
