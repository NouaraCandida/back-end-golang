package entity

import (
	"github.com/google/uuid"
)

type Sensor struct {
	ID 				uuid.UUID
	Nome            string
	Localidade 	Localidade        
}