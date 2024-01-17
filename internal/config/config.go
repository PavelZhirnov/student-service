package config

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	ServiceName string   `mapstructure:"service_name"`
	Ports       Ports    `mapstructure:"ports"`
	Postgres    Postgres `mapstructure:"postgres"`
}

type Ports struct {
	GRPC string `mapstructure:"grpc"`
}

type Postgres struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	DBName   string `mapstructure:"dbName"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

func InitConfig() (*Config, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, errors.Wrap(err, "os.Getwd")
	}
	configPath := fmt.Sprintf("%s/../../", path)

	cfg := &Config{}

	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err = viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.Wrap(err, "viper.Unmarshal")
	}

	return cfg, nil
}
