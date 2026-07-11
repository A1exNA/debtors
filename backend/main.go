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
	http.HandleFunc("/houses", handler.HousesHandler)
	http.ListenAndServe(":8080", nil)
}
