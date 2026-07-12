package handler

import (
	"debtors/dto"
	"debtors/service"
	"encoding/json"
	"net/http"
)

func (h *Handlers) CreateHouseHandler(w http.ResponseWriter, r *http.Request) {
	house := dto.HouseCreate{}

	err := json.NewDecoder(r.Body).Decode(&house)

	if err != nil {
		response := dto.Response{Status: "error", Data: err.Error()}
		writeJSON(w, response)

		return
	}

	resp, err := service.CreateHousesService(h.Conn, house.Address, house.IsServiced)

	if err != nil {
		response := dto.Response{Status: "error", Data: err.Error()}
		writeJSON(w, response)

		return
	}

	response := dto.Response{Status: "success", Data: resp}
	writeJSON(w, response)
}

func (h *Handlers) ReadHouseHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := service.ReadHousesService(h.Conn)

	if err != nil {
		response := dto.Response{Status: "error", Data: err.Error()}
		writeJSON(w, response)

		return
	}

	response := dto.Response{Status: "success", Data: resp}
	writeJSON(w, response)
}

func (h *Handlers) UpdateHouseHandler(w http.ResponseWriter, r *http.Request) {
	house := dto.House{}

	err := json.NewDecoder(r.Body).Decode(&house)

	if err != nil {
		response := dto.Response{Status: "error", Data: err.Error()}
		writeJSON(w, response)

		return
	}

	resp, err := service.UpdateHousesService(h.Conn, house.Id, house.Address, house.IsServiced)

	if err != nil {
		response := dto.Response{Status: "error", Data: err.Error()}
		writeJSON(w, response)

		return
	}

	response := dto.Response{Status: "success", Data: resp}
	writeJSON(w, response)
}

func (h *Handlers) DeleteHouseHandler(w http.ResponseWriter, r *http.Request) {
	house := dto.HouseDelete{}

	err := json.NewDecoder(r.Body).Decode(&house)

	if err != nil {
		response := dto.Response{Status: "error", Data: err.Error()}
		writeJSON(w, response)

		return
	}

	resp, err := service.DeleteHousesService(h.Conn, house.Id)

	if err != nil {
		response := dto.Response{Status: "error", Data: err.Error()}
		writeJSON(w, response)

		return
	}

	response := dto.Response{Status: "success", Data: resp}
	writeJSON(w, response)
}
