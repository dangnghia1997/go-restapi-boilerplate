package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBConnectionString string
	ServerPort         string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("config") // You can also have a separate config directory

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{
		DBConnectionString: viper.GetString("db_connection_string"),
		ServerPort:         viper.GetString("server_port"),
	}

	return cfg, nil
}
