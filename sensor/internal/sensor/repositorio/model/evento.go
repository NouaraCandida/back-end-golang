package model

import (
	"context"

	"github.com/google/uuid"
)

type EventoRepositorio interface {
	Get(ctx context.Context, id uuid.UUID) (Evento, error)
	GetEventosToIDSensor(ctx context.Context, id_idsensor uuid.UUID) ([]Evento, error)
	Create(ctx context.Context, evento Evento) (Evento, error)
}
type Evento struct {
	ID       uuid.UUID `json:"id"`
	Valor    string    `json:"valor"`
	IDSensor uuid.UUID `json:"idSensor"`
}

func NewEvento(valor string, idSensor string) (Evento, error) {
	if valor == "" {
		return Evento{}, ErrInvalidValorEvento
	}

	uidSensor, erro := uuid.Parse(valor)
	if erro != nil {
		return Evento{}, erro
	}

	return Evento{
		Valor:    valor,
		IDSensor: uidSensor,
	}, nil
}
