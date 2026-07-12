package main

import (
	"debtors/db"
	"debtors/handler"
	"fmt"
	"net/http"
)

func main() {
	conn, err := db.Connect()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(conn)
	}

	handler := handler.Handlers{Conn: conn}

	http.HandleFunc("POST /houses", handler.CreateHouseHandler)
	http.HandleFunc("GET /houses", handler.ReadHouseHandler)
	http.HandleFunc("PUT /houses", handler.UpdateHouseHandler)
	http.HandleFunc("DELETE /houses", handler.DeleteHouseHandler)

	http.HandleFunc("POST /accounts", handler.CreateAccountHandler)
	http.HandleFunc("GET /accounts", handler.ReadAccountHandler)
	http.HandleFunc("PUT /accounts", handler.UpdateAccountHandler)
	http.HandleFunc("DELETE /accounts", handler.DeleteAccountHandler)

	http.HandleFunc("POST /debts", handler.CreateDebtHandler)
	http.HandleFunc("GET /debts", handler.ReadDebtHandler)
	http.HandleFunc("PUT /debts", handler.UpdateDebtHandler)
	http.HandleFunc("DELETE /debts", handler.DeleteDebtHandler)

	http.ListenAndServe(":8080", nil)
}
