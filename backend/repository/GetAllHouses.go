package housesRepository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type House struct {
	Id         int
	Address    string
	IsServiced bool
}

func GetAllHouses(conn *pgx.Conn) ([]House, error) {
	result, err := conn.Query(context.Background(), "SELECT * FROM houses")

	if err != nil {
		return nil, err
	}

	houses := []House{}

	for result.Next() {
		house := House{}
		result.Scan(&house.Id, &house.Address, &house.IsServiced)
		houses = append(houses, house)
	}

	return houses, nil
}
