package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/lopesboa/identity-sphere/identity/presentation/middlewares"
	"github.com/lopesboa/identity-sphere/internal/config"
	"github.com/spf13/viper"
)

func Initialize() {
	app := fiber.New(fiber.Config{
		AppName:      "Identity Sphere",
		ServerHeader: "Fiber",
	})

	logger := config.GetLogger("router")

	middlewares.InitFiberMiddlewares(app, InitializePublicRoutes, nil, logger)

	var listenIp = viper.GetString("ListenIp")
	var listenPort = viper.GetString("ListenPort")

	err := app.Listen(fmt.Sprintf("%v:%v", listenIp, listenPort))

	logger.Error(err)

}
