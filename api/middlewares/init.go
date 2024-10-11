package middlewares

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/lopesboa/identity-sphere/identity/application/types"
	"github.com/lopesboa/identity-sphere/identity/infrastructure"
)

type Logger interface {
	Info(v ...interface{})
}

type publicRoutes func(app *fiber.App)
type privateRoutes func(app *fiber.App)

func useFiberCtx(fiberCtx *fiber.Ctx) error {
	var requestId = fiberCtx.Locals("requestId")
	var ctx = context.WithValue(context.Background(), types.ContextKeyRequestId, requestId)
	fiberCtx.SetUserContext(ctx)

	return fiberCtx.Next()
}

func InitFiberMiddlewares(app *fiber.App, publicRoutes publicRoutes, privateRoutes privateRoutes, localLogger Logger) {
	app.Use(requestid.New())
	app.Use(logger.New())
	app.Use(useFiberCtx)

	identityManager := infrastructure.NewIdentityManager()

	publicRoutes(app)

	app.Use(NewJwtMiddleware(identityManager))

	localLogger.Info("fiber middleware initialized")

}
