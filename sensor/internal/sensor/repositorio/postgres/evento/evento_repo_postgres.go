package repositorio

import (
	"context"
	"database/sql"
	"sensor/internal/sensor/repositorio/model"

	"github.com/google/uuid"
)

type EventoRepositorioPostgres struct {
	db *sql.DB
}

func NewRepositorioPostgres(db *sql.DB) *EventoRepositorioPostgres {
	return &EventoRepositorioPostgres{db: db}
}

func (s *EventoRepositorioPostgres) Get(ctx context.Context, id uuid.UUID) (model.Evento, error) {
	var evento model.Evento

	if err := s.db.QueryRowContext(ctx, qEvento_Get, id).
		Scan(&evento.ID, &evento.IDSensor, &evento.Valor); err != nil {
		return model.Evento{}, err
	}

	return evento, nil

}

func (s *EventoRepositorioPostgres) GetEventosToIDSensor(ctx context.Context, id_sensor uuid.UUID) ([]model.Evento, error) {

	rows, err := s.db.QueryContext(ctx, qEvento_GetEventosToIDSensor, id_sensor)
	if err != nil {
		return []model.Evento{}, err
	}
	defer rows.Close()

	var evs []model.Evento
	for rows.Next() {
		var ev model.Evento
		if err = rows.Scan(
			&ev.ID,
			&ev.IDSensor,
			&ev.Valor,
		); err != nil {
			return []model.Evento{}, err
		}
		evs = append(evs, ev)
	}

	return evs, nil

}

func (s *EventoRepositorioPostgres) Create(ctx context.Context, evento model.Evento) (model.Evento, error) {

	err := s.db.QueryRowContext(ctx, qEvento_Create, evento.IDSensor, evento.Valor).
		Scan(&evento.ID)
	if err != nil {
		return model.Evento{}, err
	}

	return evento, nil

}
