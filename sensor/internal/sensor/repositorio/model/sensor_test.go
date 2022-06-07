package model_test

import (
	"sensor/internal/sensor/repositorio/model"
	"testing"
)

func TestSensor_NewSensor(t *testing.T) {

	type testCase struct {
		test        string
		nomeSensor  string
		pais        string
		regiao      string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Validação de nome do sensor vazio",
			nomeSensor:  "",
			pais:        "Brasil",
			regiao:      "Sul",
			expectedErr: model.ErrInvalidNomeSensor,
		},

		{
			test:        "Validação pais vazio",
			nomeSensor:  "Sensor 1",
			pais:        "",
			regiao:      "Sul",
			expectedErr: model.ErrInvalidLocSensor,
		},
		{
			test:        "Validação regiao vazio",
			nomeSensor:  "Sensor 1",
			pais:        "Brasil",
			regiao:      "",
			expectedErr: model.ErrInvalidLocSensor,
		},
		{
			test:        "Valida objeto ",
			nomeSensor:  "Sensor 1",
			pais:        "Brasil",
			regiao:      "Sul",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := model.NewSensor(tc.nomeSensor, tc.pais, tc.regiao)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

		})
	}

}
