package aggregate

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"sensor/pkg/domain/sensor/v1/entidade"
	ent "sensor/pkg/domain/sensor/v1/entidade"
	vo "sensor/pkg/domain/sensor/v1/objetovalor"
)

var (
	ErrInvalidNomeSensor = errors.New("Nome do sensor é inválido")
	ErrInvalidLocSensor = errors.New("Localidade do sensor é inválida")
	ErrInvalidEventoSensor = errors.New("Valor do evento do sensor é inválido")
)

type EventoSensor struct {
	sensor 			*ent.Sensor
	evento 			*ent.Evento
	localidade 		*vo.Localidade

}

//NewEventoSensor - contém tudo que é necessário para representar EventoSensor 
func NewEventoSensor(nomeSensor string, valorEvento string, pais string, regiao string)( EventoSensor, error){
	if nomeSensor == ""{
		return EventoSensor{}, ErrInvalidNomeSensor
	}

	if valorEvento == ""{
		return EventoSensor{}, ErrInvalidEventoSensor
	}

	if pais =="" || regiao ==""{
		 return EventoSensor{}, ErrInvalidLocSensor
	}

	
	s := &entidade.Sensor{
		Nome: nomeSensor,
		ID:	  uuid.New(),
	}
	e := &entidade.Evento{
		Valor: valorEvento,
		ID:  uuid.New(),
		IDSensor: s.ID,
	}
	l := &vo.Localidade{
		NomePais: pais,
		NomeRegiao: regiao,
		CreatedAt: time.Now(),
	}
	

	return EventoSensor{
		sensor: s,
		evento: e,
		localidade: l,
	}, nil


}
