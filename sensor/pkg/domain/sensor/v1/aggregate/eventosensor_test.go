package aggregate_test

import (
	"testing"

	"sensor/pkg/domain/sensor/v1/aggregate"
)

func TestEventoSensor_NewEventoSensor(t *testing.T) {

	type testCase struct {
		test        string
		nomeSensor  string
		valorEvento string
		pais        string
		regiao      string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Validação de nome do sensor vazio",
			nomeSensor:  "",
			valorEvento: "1000",
			pais:        "Brasil",
			regiao:      "Sul",
			expectedErr: aggregate.ErrInvalidNomeSensor,
		},
		{
			test:        "Validação de valor do evento do sensor vazio",
			nomeSensor:  "Sensor 1",
			valorEvento: "",
			pais:        "Brasil",
			regiao:      "Sul",
			expectedErr: aggregate.ErrInvalidEventoSensor,
		},
		{
			test:        "Validação pais vazio",
			nomeSensor:  "Sensor 1",
			valorEvento: "1000",
			pais:        "",
			regiao:      "Sul",
			expectedErr: aggregate.ErrInvalidLocSensor,
		},
		{
			test:        "Validação regiao vazio",
			nomeSensor:  "Sensor 1",
			valorEvento: "1000",
			pais:        "Brasil",
			regiao:      "",
			expectedErr: aggregate.ErrInvalidLocSensor,
		},
		{
			test:        "Valida objeto ",
			nomeSensor:  "Sensor 1",
			valorEvento: "1000",
			pais:        "Brasil",
			regiao:      "Sul",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewEventoSensor(tc.nomeSensor, tc.valorEvento, tc.pais, tc.regiao)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

		})
	}

}
