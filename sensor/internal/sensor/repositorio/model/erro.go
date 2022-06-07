package model

import (
	"errors"
)
var (
	// ErrInvalidNomeSensor retorna erro de nome inválido.
	ErrInvalidNomeSensor = errors.New("Nome do sensor é inválido")
	// ErrInvalidLocSensor retorna erro de localidade inválido.
	ErrInvalidLocSensor = errors.New("Localidade do sensor é inválida")
	// ErrInvalidEventoSensor retorna erro de valor inválido.
	ErrInvalidEventoSensor = errors.New("Valor do evento do sensor é inválido")
)