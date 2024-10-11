package usecases

import (
	"context"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"github.com/go-playground/validator"
	"github.com/lopesboa/identity-sphere/identity/application/types"
)

type CreateUserRequest struct {
	Username     string `validate:"required,min=3,max=15"`
	Password     string `validate:"required"`
	FirstName    string `validate:"min=1,max=30"`
	LastName     string `validate:"min=1,max=30"`
	Email        string `validate:"required,email"`
	MobileNumber string
}

type CreateUserResponse struct {
	User *gocloak.User
}

type createUserUseCase struct {
	im types.IdentityManager
}

func NewCreateUserUseCase(im types.IdentityManager) *createUserUseCase {
	return &createUserUseCase{
		im: im,
	}
}

func (uc *createUserUseCase) CreateUser(ctx context.Context, params CreateUserRequest) error {
	var validate = validator.New()
	err := validate.Struct(params)

	if err != nil {
		return err
	}

	var user = gocloak.User{
		Username:      gocloak.StringP(params.Username),
		FirstName:     gocloak.StringP(params.FirstName),
		LastName:      gocloak.StringP(params.LastName),
		Email:         gocloak.StringP(params.Email),
		EmailVerified: gocloak.BoolP(false),
		Enabled:       gocloak.BoolP(true),
		Attributes:    &map[string][]string{},
		RealmRoles:    &[]string{"viewr"},
		Credentials: &[]gocloak.CredentialRepresentation{
			{
				Value: gocloak.StringP(params.Password),
			},
		},
	}

	if strings.TrimSpace(params.MobileNumber) != "" {
		(*user.Attributes)["mobile"] = []string{params.MobileNumber}
	}

	err = uc.im.CreateUser(ctx, user)

	if err != nil {
		return err
	}

	return nil
}
