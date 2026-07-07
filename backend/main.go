package main

import (
	dbConnection "debtors/db"
	housesRepository "debtors/repository"
	"fmt"
)

func main() {
	conn, err := dbConnection.Connect()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(conn)
	}

	res, err := housesRepository.GetAllHouses(conn)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
