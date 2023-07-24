package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ServiceName string `mapstructure:"SERVICE_NAME"`
	Environment string `mapstructure:"ENV"`
	Port        int    `mapstructure:"PORT"`
	BaseUrl     string `mapstructure:"BASE_URL"`

	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`

	DBUrlCollection string `mapstructure:"DB_URL_COLLECTION"`
}

func New(path, name string) (*Config, error) {
	cfg := &Config{}

	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	fmt.Print("Config: ")
	fmt.Println(cfg)
	return cfg, err
}
