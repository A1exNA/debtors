package repository

import (
	"context"
	"debtors/dto"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func CreateHouse(conn *pgx.Conn, address string, isServiced bool) (pgconn.CommandTag, error) {
	resp, err := conn.Exec(context.Background(), "INSERT INTO houses (address, is_servised) VALUES ($1, $2)", address, isServiced)

	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return resp, nil
}

func ReadAllHouses(conn *pgx.Conn) ([]dto.House, error) {
	resp, err := conn.Query(context.Background(), "SELECT * FROM houses")

	if err != nil {
		return nil, err
	}

	houses := []dto.House{}

	for resp.Next() {
		house := dto.House{}
		resp.Scan(&house.Id, &house.Address, &house.IsServiced)
		houses = append(houses, house)
	}

	return houses, nil
}

func UpdateHouse(conn *pgx.Conn, id int, address string, isServiced bool) (pgconn.CommandTag, error) {
	resp, err := conn.Exec(context.Background(), "UPDATE houses SET address = $2, is_servised = $3 WHERE id = $1", id, address, isServiced)

	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return resp, nil
}

func DeleteHouse(conn *pgx.Conn, id int) (pgconn.CommandTag, error) {
	resp, err := conn.Exec(context.Background(), "DELETE FROM houses WHERE id = $1", id)

	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return resp, nil
}
