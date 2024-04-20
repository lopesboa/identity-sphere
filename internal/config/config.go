package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var (
	logger *Logger
)

func Init() {
	logger = GetLogger("main")

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.SetEnvPrefix("bk")
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("unable to initialize viper: %w", err))
	}

	logger.Info("viper config initialized")
}
