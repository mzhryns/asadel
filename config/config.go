package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	// ServiceName string `envconfig:"SERVICE_NAME" default:"asadel-rest-api"`
	// Environment string `envconfig:"ENV" default:"dev"`
	// Port        int    `envconfig:"PORT" default:"8080" required:"true"`
	// BaseUrl     string `envconfig:"BASE_URL" default:"http://localhost:8080/" required:"true"`

	// DBHost     string `envconfig:"DB_HOST" default:"localhost"`
	// DBPort     string `envconfig:"DB_PORT" default:"27017"`
	// DBUsername string `envconfig:"DB_USERNAME" default:"root"`
	// DBPassword string `envconfig:"DB_PASSWORD" default:"example"`
	// DBName     string `envconfig:"DB_NAME" default:"asadel"`

	// DBUrlCollection string `envconfig:"DB_URL_COLLECTION" default:"urls"`

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

func New(path string) (*Config, error) {
	cfg := &Config{}
	// envconfig.MustProcess("", &cfg)

	// return cfg
	viper.AddConfigPath(path)
	viper.SetConfigName("")
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
