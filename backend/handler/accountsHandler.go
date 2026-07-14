package handler

import (
	"debtors/dto"
	"debtors/service"
	"encoding/json"
	"net/http"
)

func (h *Handlers) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	account := dto.AccountCreate{}

	err := json.NewDecoder(r.Body).Decode(&account)

	if err != nil {
		response := dto.Response{Status: "error", Data: err.Error()}
		writeJSON(w, response)

		return
	}

	data := dto.DataAccount{
		ValueAccount: dto.ValueAccount{
			AccountNumber:  &account.AccountNumber,
			HouseId:        &account.HouseId,
			PremisesType:   account.PremisesType,
			PremisesNumber: account.PremisesNumber,
			OwnerName:      account.OwnerName,
			OwnerPhone:     account.OwnerPhone,
		},
		Valid: true,
		Text:  "",
	}

	err = ValidateDataAccount(data)

	if err != nil {
		response := dto.Response{Status: "error", Data: err.Error()}
		writeJSON(w, response)

		return
	}

	resp, err := service.CreateAccountService(h.Conn, account, data)

	if err != nil {
		response := dto.Response{Status: "error", Data: err.Error()}
		writeJSON(w, response)

		return
	}

	response := dto.Response{Status: "success", Data: resp}
	writeJSON(w, response)
}

func (h *Handlers) ReadAccountHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := service.ReadAccountService(h.Conn)

	if err != nil {
		response := dto.Response{Status: "error", Data: err.Error()}
		writeJSON(w, response)

		return
	}

	response := dto.Response{Status: "success", Data: resp}
	writeJSON(w, response)
}

func (h *Handlers) UpdateAccountHandler(w http.ResponseWriter, r *http.Request) {
	account := dto.Account{}

	err := json.NewDecoder(r.Body).Decode(&account)

	if err != nil {
		response := dto.Response{Status: "error", Data: err.Error()}
		writeJSON(w, response)

		return
	}

	data := dto.DataAccount{
		ValueAccount: dto.ValueAccount{
			Id:             &account.Id,
			AccountNumber:  &account.AccountNumber,
			HouseId:        &account.HouseId,
			PremisesType:   account.PremisesType,
			PremisesNumber: account.PremisesNumber,
			OwnerName:      account.OwnerName,
			OwnerPhone:     account.OwnerPhone,
		},
		Valid: true,
		Text:  "",
	}

	err = ValidateDataAccount(data)

	if err != nil {
		response := dto.Response{Status: "error", Data: err.Error()}
		writeJSON(w, response)

		return
	}

	resp, err := service.UpdateAccountService(h.Conn, account, data)

	if err != nil {
		response := dto.Response{Status: "error", Data: err.Error()}
		writeJSON(w, response)

		return
	}

	response := dto.Response{Status: "success", Data: resp}
	writeJSON(w, response)
}

func (h *Handlers) DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	account := dto.AccountDelete{}

	err := json.NewDecoder(r.Body).Decode(&account)

	if err != nil {
		response := dto.Response{Status: "error", Data: err.Error()}
		writeJSON(w, response)

		return
	}

	data := dto.DataAccount{
		ValueAccount: dto.ValueAccount{
			Id: &account.Id,
		},
		Valid: true,
		Text:  "",
	}

	err = ValidateDataAccount(data)

	if err != nil {
		response := dto.Response{Status: "error", Data: err.Error()}
		writeJSON(w, response)

		return
	}

	resp, err := service.DeleteAccountService(h.Conn, account, data)

	if err != nil {
		response := dto.Response{Status: "error", Data: err.Error()}
		writeJSON(w, response)

		return
	}

	response := dto.Response{Status: "success", Data: resp}
	writeJSON(w, response)
}
