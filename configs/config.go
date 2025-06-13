package conifgs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db   Dbconfig
	Auth AuthConfig
}

type Dbconfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Println("error loading .env file")
	}

	return &Config{
		Db: Dbconfig{
			Dsn: os.Getenv("DSN"),
		},

		Auth: AuthConfig{
			Secret: os.Getenv("TOKEN"),
		},
	}

}
