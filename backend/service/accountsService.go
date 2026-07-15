package service

import (
	"debtors/dto"
	"debtors/repository"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func CreateAccountService(conn *pgx.Conn, account dto.AccountCreate, data dto.DataAccount) (pgconn.CommandTag, error) {
	err := ValidateDataAccountCreate(conn, data)

	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return repository.CreateAccount(conn, account)
}

func ReadAccountService(conn *pgx.Conn) ([]dto.Account, error) {
	return repository.ReadAccounts(conn)
}

func UpdateAccountService(conn *pgx.Conn, account dto.Account, data dto.DataAccount) (pgconn.CommandTag, error) {
	err := ValidateDataAccountUpdate(conn, data)

	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return repository.UpdateAccount(conn, account)
}

func DeleteAccountService(conn *pgx.Conn, account dto.AccountDelete, data dto.DataAccount) (pgconn.CommandTag, error) {
	err := ValidateDataAccountDelete(conn, data)

	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return repository.DeleteAccount(conn, account)
}
