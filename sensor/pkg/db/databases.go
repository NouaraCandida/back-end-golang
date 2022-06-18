package db

import (
	"database/sql"
	"log"
	dbpostgres "sensor/pkg/db/postgres"
)

type DataSources struct {
	DB *sql.DB
}

func InitializeDataSources() (*DataSources, error) {
	db, err := dbpostgres.PostgresConnection()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &DataSources{DB: db}, nil
}
