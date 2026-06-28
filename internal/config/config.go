package config

import "github.com/spf13/viper"

type Config struct {
	Port     string
	LogLevel string
	DBUri    string
}

func Load() *Config {
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("DB_URI", "")

	viper.AutomaticEnv()

	return &Config{
		Port:     viper.GetString("PORT"),
		LogLevel: viper.GetString("LOG_LEVEL"),
		DBUri:    viper.GetString("DB_URI"),
	}
}