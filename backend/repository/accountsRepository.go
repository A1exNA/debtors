package repository

import (
	"context"
	"debtors/dto"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func CreateAccount(conn *pgx.Conn, account dto.AccountCreate) (pgconn.CommandTag, error) {
	resp, err := conn.Exec(context.Background(), "INSERT INTO accounts (number, house_id, premises_type, premises_number) VALUES ($1, $2, $3, $4)", account.Number, account.HouseId, account.PremisesType, account.PremisesNumber)

	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return resp, nil
}

func ReadAccounts(conn *pgx.Conn) ([]dto.Account, error) {
	resp, err := conn.Query(context.Background(), "SELECT * FROM accounts ORDER BY number ASC")

	if err != nil {
		return nil, err
	}

	accounts := []dto.Account{}

	for resp.Next() {
		account := dto.Account{}
		resp.Scan(&account.Id, &account.Number, &account.HouseId, &account.PremisesType, &account.PremisesNumber)
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func UpdateAccount(conn *pgx.Conn, account dto.Account) (pgconn.CommandTag, error) {
	resp, err := conn.Exec(context.Background(), "UPDATE accounts SET number = $2, house_id = $3, premises_type = $4, premises_number = $5 WHERE id = $1", account.Id, account.Number, account.HouseId, account.PremisesType, account.PremisesNumber)

	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return resp, nil
}

func DeleteAccount(conn *pgx.Conn, account dto.AccountDelete) (pgconn.CommandTag, error) {
	resp, err := conn.Exec(context.Background(), "DELETE FROM accounts WHERE id = $1", account.Id)

	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return resp, nil
}
