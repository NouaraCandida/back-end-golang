package repositorio

import (
	"context"
	"database/sql"
	"sensor/internal/sensor/repositorio/model"
	"github.com/google/uuid"
	
)

type SensorRepositorioPostgres struct {
	db *sql.DB
}

func NewRepositorioPostgres(db *sql.DB) *SensorRepositorioPostgres {
	return &SensorRepositorioPostgres{db: db}
}

func (s *SensorRepositorioPostgres) Get(id uuid.UUID) (model.Sensor, error) {
	var sensor model.Sensor
	query := "select id, nome from sensores where id=$1"
	
	if err  :=  s.db.QueryRowContext(context.Background(),query,id).
		Scan(&sensor.ID, &sensor.Nome); err != nil{
			return model.Sensor{}, err
		}

	return sensor, nil

}

func (s *SensorRepositorioPostgres) Create(sensor model.Sensor) (model.Sensor, error) {
	query := "insert into sensores (nome, nome_regiao, nome_pais) values ($1, $2, $3) returning id"
	
	 err  :=  s.db.QueryRowContext(context.Background(),query,sensor.Nome, sensor.Nomepais, sensor.Nomeregiao).
		Scan(&sensor.ID);
	 if err != nil{
		return model.Sensor{}, err
	}

	return sensor, nil

}
