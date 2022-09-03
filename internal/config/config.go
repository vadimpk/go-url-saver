package config

import (
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	ServerConfig
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

func Init(configPath string) (*Config, error) {
	if err := parseConfigPath(configPath); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.UnmarshalKey("server", &cfg.ServerConfig); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func parseConfigPath(filepath string) error {
	path := strings.Split(filepath, "/")

	viper.AddConfigPath(path[0])
	viper.SetConfigName(path[1])

	return viper.ReadInConfig()
}
