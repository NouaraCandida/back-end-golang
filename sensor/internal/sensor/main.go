package main

import (
	"fmt"
	"log"
	db "sensor/db/postgres"
	"sensor/internal/sensor/repositorio/model"
	repo "sensor/internal/sensor/repositorio/postgres"
	"sensor/internal/sensor/service"

	"github.com/google/uuid"
)

func main() {

	db, err := db.Conectar()
	if err != nil {
		log.Fatal(err)
	}
	repositorio := repo.NewRepositorioPostgres(db)

	s := service.NewSensorService(repositorio)

	ids := "4fc8b12a-2586-4639-b7f0-4806ab681eba"
	id, err := uuid.Parse(ids)
	sensor, err := s.FindById(id)

	screate := model.Sensor{
		Nome:       "Sensor 2",
		Nomeregiao: "Norte",
		Nomepais:   "Brasil",
	}
	sss, err := s.Create(screate)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sensor)
	fmt.Println(sss)

}
