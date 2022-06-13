package mocks

import (
	"context"
	"sensor/internal/sensor/repositorio/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	
)

type MockSensorRepositorioPostgres struct {
	mock.Mock
}


func (m *MockSensorRepositorioPostgres) Get(ctx context.Context, id uuid.UUID) (model.Sensor, error) {
	ret := m.Called(ctx, id)

	var r0 model.Sensor
	if ret.Get(0) != nil{
		r0 = ret.Get(0).(model.Sensor)
	}

 	var r1 error

    if ret.Get(1) != nil {
        r1 = ret.Get(1).(error)
    }

	return r0, r1

}

func (m *MockSensorRepositorioPostgres) Create(ctx context.Context, sensor model.Sensor) (model.Sensor, error) {
	ret := m.Called(ctx, sensor)

	var r0 model.Sensor
	if ret.Get(0) != nil{
		r0 = ret.Get(0).(model.Sensor)
	}

 	var r1 error

    if ret.Get(1) != nil {
        r1 = ret.Get(1).(error)
    }

	return r0, r1

}
