package controllers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sensor/internal/sensor/repository/model"

	"sensor/pkg/config/sensor/responses"

	"github.com/google/uuid"
)

func (h *Handler) CreateSensor(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var sensor model.Sensor
	if err = json.Unmarshal(bodyRequest, &sensor); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = sensor.Validate(model.Create); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	h.SensorService.Create(context.Background(), &sensor)

	responses.JSON(w, http.StatusCreated, sensor)

}

func (h *Handler) GetSensor(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	id, err := uuid.Parse(string(bodyRequest))
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	sensor, err := h.SensorService.Get(context.Background(), id)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, sensor)
}
