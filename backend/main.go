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
	http.ListenAndServe(":8080", nil)
}
