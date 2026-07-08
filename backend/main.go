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

	resG, err := housesRepository.GetAllHouses(conn)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resG)
	}

	// resC, err := housesRepository.CreateHouse(conn, "Тестовый дом, Test house", true)
	//
	//	if err != nil {
	//		fmt.Println(err)
	//	} else {
	//
	//		fmt.Println(resC)
	//	}
}
