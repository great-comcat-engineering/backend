package configs

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App struct {
		Host       string `envconfig:"APP_HOST" default:"localhost"`
		Port       string `envconfig:"APP_PORT" default:":8080"`
		ApiVersion string `envconfig:"API_VERSION" default:"v0"`
		AppVersion string `envconfig:"APP_VERSION" default:"v0.0.1"`
	}
	Database struct {
		MongoUri      string `envconfig:"MONGODB_URI" default:"mongodb://localhost:27017"`
		AccessTimeout int    `envconfig:"MONGODB_ACCESS_TIMEOUT" default:"5"`
	}
	Auth struct {
		JWTSecret        string `envconfig:"JWT_SECRET" default:"token-secret"`
		JWTExpireInHours int    `envconfig:"JWT_EXPIRE" default:"24"`
		TokenExpire      int    `envconfig:"TOKEN_EXPIRE" default:"60"`
		ShortTokenExpire int    `envconfig:"SHORT_TOKEN_EXPIRE" default:"15"`
	}
}

var appConfig = &Config{}

func AppConfig() *Config {
	return appConfig
}

func LoadConfig() error {
	godotenv.Load()
	if err := envconfig.Process("", appConfig); err != nil {
		return err
	}

	return nil
}
