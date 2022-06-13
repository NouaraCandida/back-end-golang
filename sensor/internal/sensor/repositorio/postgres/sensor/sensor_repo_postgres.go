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

func (s *SensorRepositorioPostgres) Get(ctx context.Context, id uuid.UUID) (model.Sensor, error) {
	var sensor model.Sensor

	if err := s.db.QueryRowContext(ctx, qSensor_Get, id).
		Scan(&sensor.ID, &sensor.Nome); err != nil {
		return model.Sensor{}, err
	}

	return sensor, nil

}

func (s *SensorRepositorioPostgres) Create(ctx context.Context, sensor model.Sensor) (model.Sensor, error) {
	err := s.db.QueryRowContext(ctx, qSensor_Create, sensor.Nome, sensor.Nomepais, sensor.Nomeregiao).
		Scan(&sensor.ID)
	if err != nil {
		return model.Sensor{}, err
	}

	return sensor, nil

}
