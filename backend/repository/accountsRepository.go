package repository

import (
	"context"
	"debtors/dto"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func CreateAccount(conn *pgx.Conn, account dto.AccountCreate) (pgconn.CommandTag, error) {
	resp, err := conn.Exec(context.Background(), "INSERT INTO accounts (account_number, house_id, premises_type, premises_number, owner_name, owner_phone) VALUES ($1, $2, $3, $4, $5, $6)", account.AccountNumber, account.HouseId, account.PremisesType, account.PremisesNumber, account.OwnerName, account.OwnerPhone)

	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return resp, nil
}

func ReadAccounts(conn *pgx.Conn) ([]dto.Account, error) {
	resp, err := conn.Query(context.Background(), "SELECT * FROM accounts ORDER BY account_number ASC")

	if err != nil {
		return nil, err
	}

	accounts := []dto.Account{}

	for resp.Next() {
		account := dto.Account{}
		resp.Scan(&account.Id, &account.AccountNumber, &account.HouseId, &account.PremisesType, &account.PremisesNumber, &account.OwnerName, &account.OwnerPhone)
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func UpdateAccount(conn *pgx.Conn, account dto.Account) (pgconn.CommandTag, error) {
	resp, err := conn.Exec(context.Background(), "UPDATE accounts SET account_number = $2, house_id = $3, premises_type = $4, premises_number = $5, owner_name = $6, owner_phone = $7 WHERE id = $1", account.Id, account.AccountNumber, account.HouseId, account.PremisesType, account.PremisesNumber, account.OwnerName, account.OwnerPhone)

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
