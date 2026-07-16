package main

import (
	"debtors/db"
	"debtors/handler"
	"fmt"
	"net/http"
)

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func main() {
	conn, err := db.Connect()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(conn)
	}

	handler := handler.Handlers{Conn: conn}

	http.HandleFunc("POST /houses", enableCORS(handler.CreateHouseHandler))
	http.HandleFunc("GET /houses", enableCORS(handler.ReadHouseHandler))
	http.HandleFunc("PUT /houses", enableCORS(handler.UpdateHouseHandler))
	http.HandleFunc("DELETE /houses", enableCORS(handler.DeleteHouseHandler))
	http.HandleFunc("OPTIONS /houses", enableCORS(func(w http.ResponseWriter, r *http.Request) {}))

	http.HandleFunc("POST /accounts", enableCORS(handler.CreateAccountHandler))
	http.HandleFunc("GET /accounts", enableCORS(handler.ReadAccountHandler))
	http.HandleFunc("PUT /accounts", enableCORS(handler.UpdateAccountHandler))
	http.HandleFunc("DELETE /accounts", enableCORS(handler.DeleteAccountHandler))
	http.HandleFunc("OPTIONS /accounts", enableCORS(func(w http.ResponseWriter, r *http.Request) {}))

	http.HandleFunc("POST /debts", enableCORS(handler.CreateDebtHandler))
	http.HandleFunc("GET /debts", enableCORS(handler.ReadDebtHandler))
	http.HandleFunc("PUT /debts", enableCORS(handler.UpdateDebtHandler))
	http.HandleFunc("DELETE /debts", enableCORS(handler.DeleteDebtHandler))
	http.HandleFunc("OPTIONS /debts", enableCORS(func(w http.ResponseWriter, r *http.Request) {}))

	http.ListenAndServe(":8080", nil)
}
