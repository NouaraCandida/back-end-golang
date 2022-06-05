package entidade

import (
	"github.com/google/uuid"
)

type Evento struct {
	ID 					uuid.UUID
	Valor           	string
	IDSensor 			uuid.UUID

} 