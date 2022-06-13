package service

import (
	"context"
	"sensor/internal/sensor/repositorio/model"

	"github.com/google/uuid"
)

//DI IC - não dizemos qual repo e sim uma interface, dando abertura para mudanças
//de banco de dados sem grandes refatoração
type SensorService struct {
	Repositorio   model.SensorRepositorio
}

func NewSensorService(repositorio model.SensorRepositorio) *SensorService {
	return &SensorService{Repositorio: repositorio}
}

func (s *SensorService) Get(ctx context.Context, id string) (model.Sensor, error) {
	iduui, err := uuid.Parse(id)
	if err != nil {
		return model.Sensor{}, err
	}
	sensor, err := s.Repositorio.Get(ctx, iduui)
	if err != nil {
		return model.Sensor{}, err
	}
	return sensor, nil

}

func (s *SensorService) Create(ctx context.Context, sensor model.Sensor) (model.Sensor, error) {
	sensor, err := s.Repositorio.Create(ctx, sensor)
	if err != nil {
		return model.Sensor{}, err
	}
	return sensor, nil

}


