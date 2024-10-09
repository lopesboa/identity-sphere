package main

import (
	"github.com/lopesboa/identity-sphere/internal/handlers/routes"
	"github.com/lopesboa/identity-sphere/internal/tools"
)

func main() {

	tools.Init()

	routes.Initialize()
}
