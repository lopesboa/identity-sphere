package middlewares

import (
	"context"
	"errors"
	"strings"

	"github.com/gofiber/fiber"
	"github.com/lopesboa/identity-sphere/internal/config"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	logger := config.GetLogger("error handler")

	if err != nil {

		switch {
		case errors.Is(err, context.Canceled):
			logger.Error("Request canceled", err)
			return ctx.Status(fiber.StatusRequestTimeout).JSON(fiber.Map{"error": "Request timeout"})
		case strings.Contains(err.Error(), "409"):
			return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "User exists with same email"})
		case strings.Contains(err.Error(), "400"):
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad request"})
		case strings.Contains(err.Error(), "403"):
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Forbidden"})
		case strings.Contains(err.Error(), "406"):
			return ctx.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"error": "Not acceptable"})
		case strings.Contains(err.Error(), "422"):
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": "Unprocessable entity"})
		case strings.Contains(err.Error(), "500"):
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
		case strings.Contains(err.Error(), "501"):
			return ctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{"error": "Not implemented"})
		case strings.Contains(err.Error(), "502"):
			return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": "Bad gateway"})
		case strings.Contains(err.Error(), "503"):
			return ctx.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "Service unavailable"})
		case strings.Contains(err.Error(), "504"):
			return ctx.Status(fiber.StatusGatewayTimeout).JSON(fiber.Map{"error": "Gateway timeout"})
		default:
			logger.Error("Internal server error", err)
			return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": "Opps, somenthing want wrong"})
		}
	}
	return nil

}
