package repository

import (
	"context"
	"debtors/dto"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func CreateDebt(conn *pgx.Conn, debt dto.DebtCreate) (pgconn.CommandTag, error) {
	resp, err := conn.Exec(context.Background(), "INSERT INTO debts (account_number, report_date, opening_balance, accrued, paid, closing_balance, upload_date) VALUES ($1, $2, $3, $4, $5, $6, $7)", debt.AccountNumber, debt.ReportDate, debt.OpeningBalance, debt.Accrued, debt.Paid, debt.ClosingBalance, debt.UploadDate)

	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return resp, nil
}

func ReadDebts(conn *pgx.Conn) ([]dto.Debt, error) {
	resp, err := conn.Query(context.Background(), "SELECT * FROM debts ORDER BY id ASC")

	if err != nil {
		return nil, err
	}

	debts := []dto.Debt{}

	for resp.Next() {
		debt := dto.Debt{}
		resp.Scan(&debt.Id, &debt.AccountNumber, &debt.ReportDate, &debt.OpeningBalance, &debt.Accrued, &debt.Paid, &debt.ClosingBalance, &debt.UploadDate)
		debts = append(debts, debt)
	}

	return debts, nil
}

func UpdateDebt(conn *pgx.Conn, debt dto.Debt) (pgconn.CommandTag, error) {
	resp, err := conn.Exec(context.Background(), "UPDATE debts SET account_number = $2, report_date = $3, opening_balance = $4, accrued = $5, paid = $6, closing_balance = $7 WHERE id = $1", debt.Id, debt.AccountNumber, debt.ReportDate, debt.OpeningBalance, debt.Accrued, debt.Paid, debt.ClosingBalance)

	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return resp, nil
}

func DeleteDebt(conn *pgx.Conn, debt dto.DebtDelete) (pgconn.CommandTag, error) {
	resp, err := conn.Exec(context.Background(), "DELETE FROM debts WHERE id = $1", debt.Id)

	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return resp, nil
}
