package main

import (
	router "github.com/lopesboa/identity-sphere/identity/presentation/routes"
	"github.com/lopesboa/identity-sphere/internal/config"
)

func main() {

	config.Init()

	router.Initialize()
}
