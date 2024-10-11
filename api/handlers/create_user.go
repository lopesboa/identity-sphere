package handlers

import (
	"context"
	"strings"

	"github.com/gofiber/fiber/v2"
	usecases "github.com/lopesboa/identity-sphere/identity/application/use-cases"
	"github.com/pkg/errors"
)

type createUserUseCase interface {
	CreateUser(context.Context, usecases.CreateUserRequest) error
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
			return handleError(logger, err, "unbale to parse incoming request")
		}

		err = uc.CreateUser(ctx, request)

		if err != nil {
			return handleError(logger, err, "failed to create user")
		}

		sucessResponse := fiber.Map{
			"Status":  "success",
			"Message": "User created successfully",
		}

		return fiberCtx.Status(fiber.StatusCreated).JSON(sucessResponse)

	}
}

func handleError(logger logger, err error, message string) error {

	switch {
	case errors.Is(err, context.Canceled):
		logger.Error(message, err)
		// return ctx.Status(fiber.StatusRequestTimeout).JSON(fiber.Map{"error": "request timeout"})
		return NewApiError(fiber.StatusRequestTimeout, err)
	case strings.Contains(err.Error(), "409"):
		logger.Error(message, err)
		logger.Error(message, err)
		return NewApiError(fiber.StatusConflict, err)

		// return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "User exists with same email"})
	case strings.Contains(err.Error(), "404"):
		logger.Error(message, err)
		return NewApiError(fiber.StatusNotFound, err)

		// return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	default:
		logger.Error(message, err)
		return NewApiError(fiber.StatusBadGateway, err)
		// return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": message})
	}

}
