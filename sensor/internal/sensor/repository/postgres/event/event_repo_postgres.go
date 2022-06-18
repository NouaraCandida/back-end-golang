package repository

import (
	"context"
	"database/sql"
	"sensor/internal/sensor/repository/model"

	"github.com/google/uuid"
)

type EventRepositoryPostgres struct {
	db *sql.DB
}

func NewRepositoryPostgres(db *sql.DB) *EventRepositoryPostgres {
	return &EventRepositoryPostgres{db: db}
}

func (s *EventRepositoryPostgres) Get(ctx context.Context, id uuid.UUID) (*model.Event, error) {
	var evento model.Event

	if err := s.db.QueryRowContext(ctx, qEvento_Get, id).
		Scan(&evento.ID, &evento.IDSensor, &evento.Valor); err != nil {
		return &model.Event{}, err
	}

	return &evento, nil

}

func (s *EventRepositoryPostgres) GetEventsToIDSensor(ctx context.Context, id_sensor uuid.UUID) ([]*model.Event, error) {

	rows, err := s.db.QueryContext(ctx, qEvento_GetEventosToIDSensor, id_sensor)
	if err != nil {
		return []*model.Event{}, err
	}
	defer rows.Close()

	var evs []*model.Event
	for rows.Next() {
		var ev model.Event
		if err = rows.Scan(
			&ev.ID,
			&ev.IDSensor,
			&ev.Valor,
		); err != nil {
			return []*model.Event{}, err
		}
		evs = append(evs, &ev)
	}

	return evs, nil

}

func (s *EventRepositoryPostgres) Create(ctx context.Context, evento *model.Event) (uuid.UUID, error) {

	err := s.db.QueryRowContext(ctx, qEvento_Create, evento.IDSensor, evento.Valor).
		Scan(&evento.ID)
	if err != nil {
		return uuid.Nil, err
	}

	return evento.ID, nil

}
