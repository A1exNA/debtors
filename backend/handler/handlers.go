package handler

import (
	"debtors/dto"
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type Handlers struct {
	Conn *pgx.Conn
}

func writeJSON(w http.ResponseWriter, response dto.Response) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
