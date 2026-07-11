package service

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"debtors/dto"
	"debtors/repository"
)

func GetHousesService(conn *pgx.Conn) ([]dto.House, error) {
	return repository.GetAllHouses(conn)
}

func CreateHousesService(conn *pgx.Conn, address string, isServiced bool) (pgconn.CommandTag, error) {
	return repository.CreateHouse(conn, address, isServiced)
}

func UpdateHousesService(conn *pgx.Conn, id int, address string, isServiced bool) (pgconn.CommandTag, error) {
	return repository.UpdateHouse(conn, id, address, isServiced)
}

func DeleteHousesService(conn *pgx.Conn, id int) (pgconn.CommandTag, error) {
	return repository.DeleteHouse(conn, id)
}
