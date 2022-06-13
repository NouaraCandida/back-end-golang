package service

import (
	"context"
	"sensor/internal/sensor/repositorio/model"

	"github.com/google/uuid"
)

//DI IC - não dizemos qual repo e sim uma interface, dando abertura para mudanças
//de banco de dados sem grandes refatoração
// EventoSensorService fica todas as dependências para esse serviço
type EventoSensorService struct {
	RepositorioSensor model.SensorRepositorio
	RepositorioEvento model.EventoRepositorio
}

// EventoSensorConfig responsável pela injetar todas as interfaces necessárias no EventoSensorService
type EventoSensorConfig struct {
	RepositorioSensor model.SensorRepositorio
	RepositorioEvento model.EventoRepositorio
}

// NewEventoSensorService
func NewEventoSensorService(c *EventoSensorConfig) *EventoSensorService {
	return &EventoSensorService{
		RepositorioSensor: c.RepositorioSensor,
		RepositorioEvento: c.RepositorioEvento}
}

func (s *EventoSensorService) GetSensorEventos(ctx context.Context, id_sensor uuid.UUID) (model.Sensor, []model.Evento, error) {
	sensor, err := s.RepositorioSensor.Get(ctx, id_sensor)
	if err != nil {
		return model.Sensor{}, []model.Evento{}, err
	}
	eventos, err := s.RepositorioEvento.GetEventosToIDSensor(ctx, id_sensor)
	if err != nil {
		return model.Sensor{}, []model.Evento{}, err
	}

	return sensor, eventos, nil

}
