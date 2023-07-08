package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	ServiceName string `envconfig:"SERVICE_NAME" default:"asadel-rest-api"`
	Environment string `envconfig:"ENV" default:"dev"`
	Port        int    `envconfig:"PORT" default:"8080" required:"true"`

	DBHost     string `envconfig:"DB_HOST" default:"localhost"`
	DBPort     string `envconfig:"DB_PORT" default:"27017"`
	DBUsername string `envconfig:"DB_USERNAME" default:"root"`
	DBPassword string `envconfig:"DB_PASSWORD" default:"example"`
	DBName     string `envconfig:"DB_NAME" default:"asadel"`

	DBUrlCollection string `envconfig:"DB_URL_COLLECTION" default:"urls"`
}

func New() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)

	return cfg
}
