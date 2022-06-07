package model

import (
	"context"

	"github.com/google/uuid"
)

type SensorRepositorio interface {
	Get(ctx context.Context, id uuid.UUID) (Sensor, error)
	Create(ctx context.Context, sensor Sensor) (Sensor, error)
}

type Sensor struct {
	ID         uuid.UUID `json:"id"`
	Nome       string    `json:"nome"`
	Nomeregiao string    `json:"nomeregiao"`
	Nomepais   string    `json:"nomepais"`
}

func NewSensor(nome string, pais string, regiao string) (Sensor, error) {
	if nome == "" {
		return Sensor{}, ErrInvalidNomeSensor
	}

	if pais == "" || regiao == "" {
		return Sensor{}, ErrInvalidLocSensor
	}

	return Sensor{
		Nome:       nome,
		Nomeregiao: regiao,
		Nomepais:   pais,
	}, nil

}
