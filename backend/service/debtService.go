package service

import (
	"debtors/dto"
	"debtors/repository"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func CreateDebtService(conn *pgx.Conn, debt dto.DebtCreate) (pgconn.CommandTag, error) {
	return repository.CreateDebt(conn, debt)
}

func ReadDebtService(conn *pgx.Conn) ([]dto.Debt, error) {
	return repository.ReadDebt(conn)
}

func UpdateDebtService(conn *pgx.Conn, debt dto.Debt) (pgconn.CommandTag, error) {
	return repository.UpdateDebt(conn, debt)
}

func DeleteDebtService(conn *pgx.Conn, debt dto.DebtDelete) (pgconn.CommandTag, error) {
	return repository.DeleteDebt(conn, debt)
}
