package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	usecases "github.com/lopesboa/identity-sphere/identity/application/use-cases"
	"github.com/pkg/errors"
)

type createUserUseCase interface {
	CreateUser(context.Context, usecases.CreateUserRequest) (*usecases.CreateUserResponse, error)
}

type logger interface {
	Error(v ...interface{})
}

func HandlerCreateUser(logger logger, uc createUserUseCase) fiber.Handler {
	return func(fiberCtx *fiber.Ctx) error {
		var ctx = fiberCtx.UserContext()
		var request = usecases.CreateUserRequest{}

		err := fiberCtx.BodyParser(&request)

		if err != nil {
			return handleError(fiberCtx, logger, err, "unbale to parse incoming request")
		}

		response, err := uc.CreateUser(ctx, request)

		if err != nil {
			return handleError(fiberCtx, logger, err, "failed to create user")
		}

		return fiberCtx.Status(fiber.StatusCreated).JSON(response)

	}
}

func handleError(ctx *fiber.Ctx, logger logger, err error, message string) error {
	if errors.Is(err, context.Canceled) {
		logger.Error(message, err)
		return ctx.Status(fiber.StatusRequestTimeout).JSON(fiber.Map{"error": "request timeout"})
	}

	logger.Error(message, err)
	return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": message})
}
