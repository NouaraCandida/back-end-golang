package repositorio

import (
	"context"


	"github.com/google/uuid"
)

type EventoSensor interface {

	GetEventoSensor(ctx context.Context, ID uuid.UUID) (*EventoSensor, error)
	SaveEventoSensor(ctx context.Context, evento EventoSensor) (*EventoSensor, error)

}