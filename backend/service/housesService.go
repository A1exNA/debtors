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

func UpdateHouseService(conn *pgx.Conn, house dto.House) (pgconn.CommandTag, error) {
	return repository.UpdateHouse(conn, house)
}

func DeleteHouseService(conn *pgx.Conn, house dto.HouseDelete) (pgconn.CommandTag, error) {
	return repository.DeleteHouse(conn, house)
}
