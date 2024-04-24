package middlewares

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/lopesboa/identity-sphere/identity/application/types"
)

type Logger interface {
	Info(v ...interface{})
}

func useFiberCtx(fiberCtx *fiber.Ctx) error {
	var requestId = fiberCtx.Locals("requestId")
	var ctx = context.WithValue(context.Background(), types.ContextKeyRequestId, requestId)
	fiberCtx.SetUserContext(ctx)

	return fiberCtx.Next()
}

func InitFiberMiddlewares(app *fiber.App, publicRoutes func(app *fiber.App), privateRoutes func(app *fiber.App), l Logger) {
	app.Use(requestid.New())
	app.Use(logger.New())
	app.Use(useFiberCtx)

	publicRoutes(app)

	l.Info("fiber middleware initialized")

}
