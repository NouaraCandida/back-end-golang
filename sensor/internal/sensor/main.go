package main

import (
	"context"
	"fmt"
	"log"
	db "sensor/db/postgres"

	"sensor/internal/sensor/repositorio/model"
	repoEvento "sensor/internal/sensor/repositorio/postgres/evento"
	repo "sensor/internal/sensor/repositorio/postgres/sensor"

	"github.com/google/uuid"
)

func main() {

	/*obj, err := model.NewEvento("valor", "ad4c3174-fd37-42af-ada4-ceb2f457d457")
	fmt.Print(err)
	fmt.Print(obj)*/

	// A mudança de repositorio aconteceria somente até repositorio, pois NewSensorService tem como parÂmetro uma interface
	/*db, err := db.Conectar()
	if err != nil {
		log.Fatal(err)
	}
	repositorio := repo.NewRepositorioPostgres(db)

	s := service.NewSensorService(repositorio)

	ids := "4fc8b12a-2586-4639-b7f0-4806ab681eba"
	id, err := uuid.Parse(ids)
	sensor, err := s.FindById(context.Background(),id)

	screate := model.Sensor{
		Nome:       "Sensor 2",
		Nomeregiao: "Norte",
		Nomepais:   "Brasil",
	}
	sss, err := s.Create(context.Background(), screate)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sensor)
	fmt.Println(sss)*/

	db, err := db.Conectar()
	if err != nil {
		log.Fatal(err)
	}

	id_sensor_test, _ := uuid.Parse("3a0944dd-2498-4fd3-93b2-30cc224956c2")
	r2 := repoEvento.NewRepositorioPostgres(db)
	r2.Create(context.Background(), model.Evento{IDSensor: id_sensor_test, Valor: "25a"})

	repositorio := repo.NewRepositorioPostgres(db)

	id_sensor_test, _ = uuid.Parse("d021b5cb-5e21-4add-bbfc-e450935ff9aa")
	result, err := repositorio.Get(context.Background(), id_sensor_test)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

}
