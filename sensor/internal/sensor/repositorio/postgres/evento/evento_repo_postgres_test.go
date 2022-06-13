package repositorio

import (
	"context"
	"errors"
	"regexp"
	"testing"

	dbpostgres "sensor/db/postgres"
	"sensor/internal/sensor/repositorio/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

func Test_eventoRepo_Create(t *testing.T) {

	type testCase struct {
		name       string
		ctx        context.Context
		input      model.Evento
		beforeTest func(sqlmock.Sqlmock)
		want       model.Evento
		wantErr    bool
	}

	idSensor, _ := uuid.Parse("3a0944dd-2498-4fd3-93b2-30cc224956c2")

	tests := []testCase{
		{
			name:  "Sucesso ao criar evento",
			ctx:   context.TODO(),
			input: model.Evento{IDSensor: idSensor, Valor: "25A"},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta(`
						INSERT INTO eventos (id_sensor, valor)
						VALUES ($1, $2);`,
					)).
					WithArgs(idSensor, "25A").
					WillReturnError(errors.New("whoops, error"))
			},
			want:    model.Evento{IDSensor: idSensor, Valor: "25A"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			db, err := dbpostgres.Conectar()
			u := NewRepositorioPostgres(db)

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			got, err := u.Create(tt.ctx, tt.input)
			if err != nil {
				t.Errorf("eventoRepo.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Valor != tt.want.Valor {
				t.Errorf("eventoRepo.Create() = %v, want %v", got.Valor, tt.want.Valor)
			}
			if got.IDSensor != tt.want.IDSensor {
				t.Errorf("eventoRepo.Create() = %v, want %v", got.IDSensor, tt.want.IDSensor)
			}
		})
	}

}

func Test_eventoRepo_Get(t *testing.T) {

	type testCase struct {
		name       string
		ctx        context.Context
		input      uuid.UUID
		beforeTest func(sqlmock.Sqlmock)
		want       model.Evento
		wantErr    bool
	}
	id_test, _ := uuid.Parse("ec943ce3-0091-4625-87e0-b2a7b4c4bd77")
	id_sensor_test, _ := uuid.Parse("3a0944dd-2498-4fd3-93b2-30cc224956c2")
	valor_test := "20A"

	tests := []testCase{
		{
			name:  "Sucesso ao pegar evento",
			ctx:   context.TODO(),
			input: id_test,
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta(qEvento_Get,
					)).
					WithArgs(id_test).
					WillReturnError(errors.New("whoops, error"))
			},
			want:    model.Evento{ID: id_test, IDSensor: id_sensor_test, Valor: valor_test},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			db, _ := dbpostgres.Conectar()
			u := NewRepositorioPostgres(db)

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			got, err := u.Get(tt.ctx, tt.input)
			if err != nil {
				t.Errorf("eventoRepo.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.ID != tt.want.ID {
				t.Errorf("eventoRepo.Get() = %v, want %v", got.ID, tt.want.ID)
			}
			if got.IDSensor != tt.want.IDSensor {
				t.Errorf("eventoRepo.get() = %v, want %v", got.IDSensor, tt.want.IDSensor)
			}
			if got.Valor != tt.want.Valor {
				t.Errorf("eventoRepo.Get() = %v, want %v", got.Valor, tt.want.Valor)
			}

		})
	}

}



func Test_eventoRepo_GetEventosToIDSensor(t *testing.T) {

	type testCase struct {
		name       string
		ctx        context.Context
		input      uuid.UUID
		beforeTest func(sqlmock.Sqlmock)
	}
	id_sensor_test, _ := uuid.Parse("3a0944dd-2498-4fd3-93b2-30cc224956c2")

	tests := []testCase{
		{
			name:  "Sucesso ao pegar eventoS",
			ctx:   context.TODO(),
			input: id_sensor_test,
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta(qEvento_GetEventosToIDSensor,
					)).
					WithArgs(id_sensor_test).
					WillReturnError(errors.New("whoops, error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			db, _ := dbpostgres.Conectar()
			u := NewRepositorioPostgres(db)

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			got, err := u.GetEventosToIDSensor(tt.ctx, tt.input)
			if err != nil {
				t.Errorf("eventoRepo.GetEventosToIDSensor() error = %v", err)
				return
			}
			if len(got) < 0{
				t.Errorf("eventoRepo.GetEventosToIDSensor() error = %v", err)
				return
			}
			

		})
	}

}