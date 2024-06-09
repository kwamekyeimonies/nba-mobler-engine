package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Keystore struct {
	DBConnectionURL string
	ServerAddress   string
}

func LoadCredentials() *Keystore {

	err := godotenv.Load("./config/appKeys.env")
	if err != nil {
		log.Fatal("unable to fetch credentials")
		return nil
	}
	cfg := &Keystore{
		DBConnectionURL: os.Getenv("POSTGRES_DB_URL"),
		ServerAddress:   os.Getenv("SERVER_ADDRESS"),
	}

	return cfg

}

func NewKeystoreConfig() *Keystore {
	keys := LoadCredentials()

	return keys

}
