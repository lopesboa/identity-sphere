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

func (uc *createUserUseCase) CreateUser(ctx context.Context, request CreateUserRequest) (*CreateUserResponse, error) {
	var validate = validator.New()
	err := validate.Struct(request)

	if err != nil {
		return nil, err
	}

	var user = gocloak.User{
		Username:      gocloak.StringP(request.Username),
		FirstName:     gocloak.StringP(request.FirstName),
		LastName:      gocloak.StringP(request.LastName),
		Email:         gocloak.StringP(request.Email),
		EmailVerified: gocloak.BoolP(false),
		Enabled:       gocloak.BoolP(true),
		Attributes:    &map[string][]string{},
		ClientRoles:   &map[string][]string{},
	}

	if strings.TrimSpace(request.MobileNumber) != "" {
		(*user.Attributes)["mobile"] = []string{request.MobileNumber}
	}
	(*user.ClientRoles)["bakongo-backend"] = []string{"car:read"}

	userResponse, err := uc.im.CreateUser(ctx, user, request.Password)

	if err != nil {
		return nil, err
	}

	var response = &CreateUserResponse{User: userResponse}

	return response, nil
}
