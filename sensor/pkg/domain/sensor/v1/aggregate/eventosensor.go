package aggregate

import (
	"errors"

	"sensor/pkg/domain/sensor/v1/model"
)

var (
	ErrInvalidNomeSensor = errors.New("Nome do sensor é inválido")
	ErrInvalidLocSensor = errors.New("Localidade do sensor é inválida")
)

type EventoSensor struct {
	sensor *model.Sensor

}

func NewEventoSensor(nome string, localidade *model.Localidade)( EventoSensor, error){
	if nome == ""{
		return EventoSensor{}, ErrInvalidNomeSensor
	}

	if localidade == nil {
		return EventoSensor{}, ErrInvalidLocSensor
	}

	evsensor := &model.Localidade{
		N
	}


}
