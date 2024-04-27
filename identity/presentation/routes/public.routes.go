package router

import (
	"github.com/gofiber/fiber/v2"
	usecases "github.com/lopesboa/identity-sphere/identity/application/use-cases"
	"github.com/lopesboa/identity-sphere/identity/infrastructure"
	"github.com/lopesboa/identity-sphere/identity/presentation/handlers"
	"github.com/lopesboa/identity-sphere/internal/config"
)

func InitializePublicRoutes(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to Identity Sphere"))
	})

	basePath := "/api/v1"

	group := app.Group(basePath)

	im := infrastructure.NewIdentityManager()
	createUserUseCase := usecases.NewCreateUserUseCase(im)
	logger := config.GetLogger("public routes")

	group.Post("/users", handlers.HandlerCreateUser(logger, createUserUseCase))
}
