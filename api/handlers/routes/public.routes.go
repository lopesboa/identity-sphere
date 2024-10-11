package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lopesboa/identity-sphere/api/handlers"
	usecases "github.com/lopesboa/identity-sphere/identity/application/use-cases"
	"github.com/lopesboa/identity-sphere/identity/infrastructure"
	"github.com/lopesboa/identity-sphere/internal/tools"
)

func InitializePublicRoutes(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to Identity Sphere"))
	})

	basePath := "/api/v1"

	group := app.Group(basePath)

	im := infrastructure.NewIdentityManager()
	lg := infrastructure.NewLoginManager()

	createUserUseCase := usecases.NewCreateUserUseCase(im)

	logger := tools.GetLogger("public routes")

	group.Post("/users", handlers.HandlerCreateUser(logger, createUserUseCase))

	group.Post("/login", handlers.HandlerLoginHandler(logger, lg))
}
