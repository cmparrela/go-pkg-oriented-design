package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	DatabaseUser     string `envconfig:"DATABASE_USER" default:"root"`
	DatabasePassword string `envconfig:"DATABASE_PASSWORD" default:"root"`
	DatabaseName     string `envconfig:"DATABASE_NAME" default:"pkg-design-example"`
	DatabasePort     string `envconfig:"DATABASE_PORT" default:"3306"`
	MongoAddress     string `envconfig:"DATABASE_ADDRESS" default:"mongodb://localhost:27017"`
}

func SetupEnvFile() *Config {
	envConfig := &Config{}
	_ = godotenv.Load()
	err := envconfig.Process("", envConfig)
	if err != nil {
		log.Fatal(nil, "Fatal error ", err)
	}

	return envConfig
}
