package entity

import (
	"github.com/google/uuid"
)

type Localidade struct {
	ID 				uuid.UUID
	Pais            Pais
	Regiao			Regiao

}