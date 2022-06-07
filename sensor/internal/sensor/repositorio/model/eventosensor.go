package model

import (
	"github.com/google/uuid"
)

type EventoSensor struct {
	sensor *Sensor
	evento *Evento
}

//NewEventoSensor - contém tudo que é necessário para representar EventoSensor
func NewEventoSensor(nomeSensor string, valorEvento string, pais string, regiao string) (EventoSensor, error) {
	if nomeSensor == "" {
		return EventoSensor{}, ErrInvalidNomeSensor
	}

	if valorEvento == "" {
		return EventoSensor{}, ErrInvalidEventoSensor
	}

	if pais == "" || regiao == "" {
		return EventoSensor{}, ErrInvalidLocSensor
	}

	s := &Sensor{
		//Nome: nomeSensor,
		//ID:	  uuid.New(),
	}
	e := &Evento{
		Valor: valorEvento,
		ID:    uuid.New(),
		//IDSensor: s.ID,
	}

	return EventoSensor{
		sensor: s,
		evento: e,
	}, nil

}

func (es *EventoSensor) GetEventoID() uuid.UUID {
	return es.evento.ID
}

/*func (es *EventoSensor) GetSensorID() uuid.UUID{
	return es.sensor.ID
}*/

func (es *EventoSensor) GetEventoValor() string {
	return es.evento.Valor
}

//func (es *EventoSensor) GetSensorNome() string {
//	return es.sensor.Nome
//}
