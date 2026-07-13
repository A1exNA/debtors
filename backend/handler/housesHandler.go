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

	data := Data{
		Value: Value{
			Address: &house.Address,
		},
		Valid: true,
		Text:  "",
	}

	data = ValidateData(data)

	if !data.Valid {
		response := dto.Response{Status: "error", Data: data.Text}
		writeJSON(w, response)

		return
	}

	resp, err := service.CreateHouseService(h.Conn, house)

	if err != nil {
		response := dto.Response{Status: "error", Data: err.Error()}
		writeJSON(w, response)

		return
	}

	response := dto.Response{Status: "success", Data: resp}
	writeJSON(w, response)
}

func (h *Handlers) ReadHouseHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := service.ReadHouseService(h.Conn)

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

	data := Data{
		Value: Value{
			Id:      &house.Id,
			Address: &house.Address,
		},
		Valid: true,
		Text:  "",
	}

	data = ValidateData(data)

	if !data.Valid {
		response := dto.Response{Status: "error", Data: data.Text}
		writeJSON(w, response)

		return
	}

	resp, err := service.UpdateHouseService(h.Conn, house)

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

	data := Data{
		Value: Value{
			Id: &house.Id,
		},
		Valid: true,
		Text:  "",
	}

	data = ValidateData(data)

	if !data.Valid {
		response := dto.Response{Status: "error", Data: data.Text}
		writeJSON(w, response)

		return
	}

	resp, err := service.DeleteHouseService(h.Conn, house)

	if err != nil {
		response := dto.Response{Status: "error", Data: err.Error()}
		writeJSON(w, response)

		return
	}

	response := dto.Response{Status: "success", Data: resp}
	writeJSON(w, response)
}
