package handler

import (
	"debtors/dto"
	"debtors/service"
	"encoding/json"
	"net/http"
)

func (h *Handlers) CreateDebtHandler(w http.ResponseWriter, r *http.Request) {
	debt := dto.DebtCreate{}

	err := json.NewDecoder(r.Body).Decode(&debt)

	if err != nil {
		response := dto.Response{Status: "error", Data: err.Error()}
		writeJSON(w, response)

		return
	}

	data := dto.DataDebt{
		ValueDebt: dto.ValueDebt{
			AccountNumber: &debt.AccountNumber,
			ReportDate:    &debt.ReportDate,
		},
		Valid: true,
		Text:  "",
	}

	err = ValidateDataDebt(data)

	if err != nil {
		response := dto.Response{Status: "error", Data: err.Error()}
		writeJSON(w, response)

		return
	}

	resp, err := service.CreateDebtService(h.Conn, debt, data)

	if err != nil {
		response := dto.Response{Status: "error", Data: err.Error()}
		writeJSON(w, response)

		return
	}

	response := dto.Response{Status: "success", Data: resp}
	writeJSON(w, response)
}

func (h *Handlers) ReadDebtHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := service.ReadDebtService(h.Conn)

	if err != nil {
		response := dto.Response{Status: "error", Data: err.Error()}
		writeJSON(w, response)

		return
	}

	response := dto.Response{Status: "success", Data: resp}
	writeJSON(w, response)
}

func (h *Handlers) UpdateDebtHandler(w http.ResponseWriter, r *http.Request) {
	debt := dto.Debt{}

	err := json.NewDecoder(r.Body).Decode(&debt)

	if err != nil {
		response := dto.Response{Status: "error", Data: err.Error()}
		writeJSON(w, response)

		return
	}

	data := dto.DataDebt{
		ValueDebt: dto.ValueDebt{
			Id:             &debt.Id,
			AccountNumber:  &debt.AccountNumber,
			ReportDate:     &debt.ReportDate,
			OpeningBalance: &debt.OpeningBalance,
			Accrued:        &debt.Accrued,
			Paid:           &debt.Paid,
			ClosingBalance: &debt.ClosingBalance,
		},
		Valid: true,
		Text:  "",
	}

	err = ValidateDataDebt(data)

	if err != nil {
		response := dto.Response{Status: "error", Data: err.Error()}
		writeJSON(w, response)

		return
	}

	resp, err := service.UpdateDebtService(h.Conn, debt, data)

	if err != nil {
		response := dto.Response{Status: "error", Data: err.Error()}
		writeJSON(w, response)

		return
	}

	response := dto.Response{Status: "success", Data: resp}
	writeJSON(w, response)
}

func (h *Handlers) DeleteDebtHandler(w http.ResponseWriter, r *http.Request) {
	debt := dto.DebtDelete{}

	err := json.NewDecoder(r.Body).Decode(&debt)

	if err != nil {
		response := dto.Response{Status: "error", Data: err.Error()}
		writeJSON(w, response)

		return
	}

	data := dto.DataDebt{
		ValueDebt: dto.ValueDebt{
			Id: &debt.Id,
		},
		Valid: true,
		Text:  "",
	}

	err = ValidateDataDebt(data)

	if err != nil {
		response := dto.Response{Status: "error", Data: err.Error()}
		writeJSON(w, response)

		return
	}

	resp, err := service.DeleteDebtService(h.Conn, debt, data)

	if err != nil {
		response := dto.Response{Status: "error", Data: err.Error()}
		writeJSON(w, response)

		return
	}

	response := dto.Response{Status: "success", Data: resp}
	writeJSON(w, response)
}
