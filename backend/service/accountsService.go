package service

import (
	"debtors/dto"
	"debtors/repository"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func CreateAccountService(conn *pgx.Conn, account dto.AccountCreate) (pgconn.CommandTag, error) {
	return repository.CreateAccount(conn, account)
}

func ReadAccountService(conn *pgx.Conn) ([]dto.Account, error) {
	return repository.ReadAccount(conn)
}

func UpdateAccountService(conn *pgx.Conn, account dto.Account) (pgconn.CommandTag, error) {
	return repository.UpdateAccount(conn, account)
}

func DeleteAccountService(conn *pgx.Conn, account dto.AccountDelete) (pgconn.CommandTag, error) {
	return repository.DeleteAccount(conn, account)
}
