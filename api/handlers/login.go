package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/lopesboa/identity-sphere/identity/infrastructure"
)

type LoginRequest struct {
	Username string `validate:"required,min=3,max=15"`
	Password string `validate:"required"`
}

type loginClient interface {
	Login(ctx context.Context, params infrastructure.LoginParams) (*infrastructure.LoginResponse, error)
}

func HandlerLoginHandler(logger logger, lg loginClient) fiber.Handler {
	return func(fiberCtx *fiber.Ctx) error {
		ctx := fiberCtx.UserContext()

		request := LoginRequest{}

		err := fiberCtx.BodyParser(&request)

		if err != nil {
			return InvalidRequestData(map[string]string{"error": err.Error()})
		}

		logger.Error(err, request)

		response, err := lg.Login(ctx, infrastructure.LoginParams{Username: request.Username, Password: request.Password})

		if err != nil {
			return handleError(logger, err, "failed to login")
		}

		return fiberCtx.Status(fiber.StatusOK).JSON(response)
	}

}
