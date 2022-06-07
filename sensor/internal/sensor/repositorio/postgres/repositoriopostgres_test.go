package repositorio_test

import (
	"context"
	"errors"
	"regexp"
	"testing"

	dbpostgres "sensor/db/postgres"
	"sensor/internal/sensor/repositorio/model"
	repo "sensor/internal/sensor/repositorio/postgres"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_sensorRepo_Create(t *testing.T) {


	type testCase struct {
		name       string
		ctx context.Context
		input model.Sensor
		beforeTest func(sqlmock.Sqlmock)
		want       model.Sensor
		wantErr    bool
	}

	tests := []testCase {
		{
			name: "Falha ao criar sensor",
			ctx:   context.TODO(),
			input: model.Sensor{Nome: "Sensor 1", Nomeregiao: "Sudeste", Nomepais: "Brasil"},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta(`
						INSERT INTO sensores (nome, nome_regiao, nome_pais)
						VALUES ($1, $2, $3);`,
					)).
					WithArgs("Sensor 1", "Sudeste", "Brasil").
					WillReturnError(errors.New("whoops, error"))
			},
			want: model.Sensor{Nome:"Sensor 1",Nomeregiao:  "Sudeste", Nomepais: "Brasil"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()
			
			db, err := dbpostgres.Conectar()
			u := repo.NewRepositorioPostgres(db)

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			got, err := u.Create(tt.input)
			if (err != nil) {
				t.Errorf("userRepo.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if(got.Nome != tt.want.Nome){
				t.Errorf("userRepo.Create() = %v, want %v", got.Nome, tt.want.Nome)
			}
			if(got.Nomepais != tt.want.Nomepais){
				t.Errorf("userRepo.Create() = %v, want %v", got.Nome, tt.want.Nome)
			}
			if(got.Nomeregiao != tt.want.Nomeregiao){
				t.Errorf("userRepo.Create() = %v, want %v", got.Nome, tt.want.Nome)
			}
		})
	}

}