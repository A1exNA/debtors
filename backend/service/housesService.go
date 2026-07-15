package service

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"debtors/dto"
	"debtors/repository"
)

func CreateHouseService(conn *pgx.Conn, house dto.HouseCreate) (pgconn.CommandTag, error) {
	return repository.CreateHouse(conn, house)
}

func ReadHouseService(conn *pgx.Conn) ([]dto.House, error) {
	return repository.ReadHouses(conn)
}

func UpdateHouseService(conn *pgx.Conn, house dto.House, data dto.DataHouse) (pgconn.CommandTag, error) {
	err := ValidateDataHouseUpdate(conn, data)

	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return repository.UpdateHouse(conn, house)
}

func DeleteHouseService(conn *pgx.Conn, house dto.HouseDelete, data dto.DataHouse) (pgconn.CommandTag, error) {
	err := ValidateDataHouseDelete(conn, data)

	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return repository.DeleteHouse(conn, house)
}
