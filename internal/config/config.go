package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

func InitViper() {

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.SetEnvPrefix("bk")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("unable to initialize viper: %w", err))
	}

	log.Println("viper config initialized")

}
