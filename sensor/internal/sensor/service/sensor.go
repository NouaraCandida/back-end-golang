package service

import (
	"sensor/internal/sensor/repositorio/model"

	"github.com/google/uuid"
)

//DI IC - não dizemos qual repo e sim uma interface, dando abertura para mudanças
//de banco de dados sem grandes refatoração
type SensorService struct {
	Repositorio model.SensorRepositorio
}

func NewSensorService(repositorio model.SensorRepositorio) *SensorService {
	return &SensorService{Repositorio: repositorio}
}

func (s *SensorService) FindById(id uuid.UUID) (model.Sensor, error) {
	sensor, err := s.Repositorio.Get(id)
	if err != nil {
		return model.Sensor{}, err
	}
	return sensor, nil

}

func (s *SensorService) Create(sensor model.Sensor) (model.Sensor, error) {
	sensor, err := s.Repositorio.Create(sensor)
	if err != nil {
		return model.Sensor{}, err
	}
	return sensor, nil

}
