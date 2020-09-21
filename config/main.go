package config

import (
	"os"

	"github.com/spf13/viper"
)

// Init - Load Config files
func Init() {
	// Set env varible
	var env string
	if os.Getenv("ENV") != "" {
		env = os.Getenv("ENV")
	} else {
		env = "dev"
	}

	// Set Config
	viper.SetConfigType("toml")
	viper.SetConfigName(env)
	viper.AddConfigPath("data/config")
	viper.ReadInConfig()
}
