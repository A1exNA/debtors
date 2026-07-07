package main

import (
	dbConnection "debtors/db"
	"fmt"
)

func main() {
	conn, err := dbConnection.Connect()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(conn)
	}
}
