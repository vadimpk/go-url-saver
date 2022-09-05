package config

import (
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Server ServerConfig
	DB     DBConfig
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func Init(configPath string) (*Config, error) {
	if err := parseConfigPath(configPath); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.UnmarshalKey("server", &cfg.Server); err != nil {
		return nil, err
	}
	if err := viper.UnmarshalKey("db", &cfg.DB); err != nil {
		return nil, err
	}

	if err := cfg.parseEnv(); err != nil {
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

func (cfg *Config) parseEnv() error {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	cfg.DB.Password = viper.GetString("DB_PASSWORD")
	return nil
}
