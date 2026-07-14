package service

import (
	"debtors/dto"
	"debtors/repository"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func CreateDebtService(conn *pgx.Conn, debt dto.DebtCreate, data dto.DataDebt) (pgconn.CommandTag, error) {
	err := ValidateDataDebtCreate(conn, data)

	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return repository.CreateDebt(conn, debt)
}

func ReadDebtService(conn *pgx.Conn) ([]dto.Debt, error) {
	return repository.ReadDebts(conn)
}

func UpdateDebtService(conn *pgx.Conn, debt dto.Debt, data dto.DataDebt) (pgconn.CommandTag, error) {
	err := ValidateDataDebtUpdate(conn, data)

	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return repository.UpdateDebt(conn, debt)
}

func DeleteDebtService(conn *pgx.Conn, debt dto.DebtDelete, data dto.DataDebt) (pgconn.CommandTag, error) {
	err := ValidateDataDebtDelete(conn, data)

	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return repository.DeleteDebt(conn, debt)
}
