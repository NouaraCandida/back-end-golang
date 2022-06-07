package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigDatabase struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func CarregaConfigDB() (*ConfigDatabase, error) {

	if erro := godotenv.Load(); erro != nil {
		log.Fatal(erro)
		return nil, erro
	}

	var config ConfigDatabase
	config.Host = os.Getenv("DATABASE_HOST")
	config.Port = os.Getenv("DATABASE_PORT")
	config.User = os.Getenv("DATABASE_USER")
	config.Password = os.Getenv("DATABASE_PASSWORD")
	config.Database = os.Getenv("DATABASE_NAME")



	return &config, nil

}
