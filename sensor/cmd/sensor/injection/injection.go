package injection

import (
	handler "sensor/internal/sensor/port/controllers"
	repoEvent "sensor/internal/sensor/repository/postgres/event"
	repoSensor "sensor/internal/sensor/repository/postgres/sensor"
	"sensor/internal/sensor/service"
	db "sensor/pkg/db"
	"time"

	"github.com/gorilla/mux"
)

func Injection(datasources *db.DataSources) (*mux.Router, error) {

	/*
	 * Repository Layer
	 */
	repoSensor := repoSensor.NewRepositoryPostgres(datasources.DB)
	repoEvent := repoEvent.NewRepositoryPostgres(datasources.DB)

	/*
	 * Service Layer
	 */
	sensorEventService := service.NewEventoSensorService(
		&service.EventSensorConfig{RepositorySensor: repoSensor, RepositoryEvent: repoEvent})

	eventService := service.NewEventService(
		&service.EventConfig{RepositoryEvent: repoEvent})
	
	sensorService := service.NewSensorService(
		&service.SensorConfig{RepositorySensor: repoSensor})


	/*
	 * Routers
	 */
	router := mux.NewRouter()
	handler.NewHandler(&handler.HandlerConfig{
		Router:              router,
		SensorEventService: sensorEventService,
		EventService: eventService,
		SensorService: sensorService,
		MaxBodyBytes:        500,
		TimeoutDuration:     time.Duration(time.Duration(5000) * time.Second),
	})

	return router, nil

}
