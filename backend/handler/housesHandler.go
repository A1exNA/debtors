package handler

import (
	"debtors/dto"
	"debtors/service"
	"encoding/json"
	"net/http"
)

func (h *Handlers) HousesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		resp, err := service.GetHousesService(h.Conn)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)

	case "POST":
		house := dto.HouseCreate{}

		err := json.NewDecoder(r.Body).Decode(&house)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp, err := service.CreateHousesService(h.Conn, house.Address, house.IsServiced)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)

	case "PUT":
		house := dto.House{}

		err := json.NewDecoder(r.Body).Decode(&house)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp, err := service.UpdateHousesService(h.Conn, house.Id, house.Address, house.IsServiced)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)

	case "DELETE":
		house := dto.HouseDelete{}

		err := json.NewDecoder(r.Body).Decode(&house)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp, err := service.DeleteHousesService(h.Conn, house.Id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
