package service

import (
	"context"
	"sensor/internal/sensor/repository/model"

	"github.com/google/uuid"
)

//DI IC - não dizemos qual repo e sim uma interface, dando abertura para mudanças
//de banco de dados sem grandes refatoração
type SensorService struct {
	repository model.SensorRepository
}

type SensorConfig struct {
	RepositorySensor model.SensorRepository
}

func NewSensorService(c *SensorConfig) *SensorService {
	return &SensorService{repository: c.RepositorySensor}
}

func (s *SensorService) Get(ctx context.Context, id uuid.UUID) (*model.Sensor, error) {

	sensor, err := s.repository.Get(ctx, id)
	if err != nil {
		return &model.Sensor{}, err
	}
	return sensor, nil

}

func (s *SensorService) Create(ctx context.Context, sensor *model.Sensor) (uuid.UUID, error) {
	id, err := s.repository.Create(ctx, sensor)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil

}
