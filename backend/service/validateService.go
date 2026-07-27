package service

import (
	"debtors/dto"
	"debtors/repository"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5"
)

func ValidateDataHouseUpdate(conn *pgx.Conn, house dto.DataHouse) error {
	houses, err := repository.ReadHouses(conn)

	if err != nil {
		return err
	}

	validId := "Такого ID нет. "
	house.Valid = false

	for _, value := range houses {
		if value.Id != *house.Id {
			continue
		}

		validId = ""
		house.Valid = true

		if value.Address == *house.Address && value.IsServiced == *house.IsServiced {
			house.Text = house.Text + "Данные не изменились. "
			house.Valid = false
		}

		break
	}

	house.Text = strings.TrimSpace(validId + house.Text)

	if !house.Valid {
		return errors.New(house.Text)
	}

	return nil
}

func ValidateDataHouseDelete(conn *pgx.Conn, house dto.DataHouse) error {
	houses, err := repository.ReadHouses(conn)

	if err != nil {
		return err
	}

	validId := "Такого ID нет. "
	house.Valid = false

	for _, value := range houses {
		if value.Id != *house.Id {
			continue
		}

		validId = ""
		house.Valid = true

		break
	}

	house.Text = strings.TrimSpace(validId + house.Text)

	if !house.Valid {
		return errors.New(house.Text)
	}

	return nil
}

func ValidateDataAccountCreate(conn *pgx.Conn, account dto.DataAccount) error {
	houses, err := repository.ReadHouses(conn)

	if err != nil {
		return err
	}

	accounts, err := repository.ReadAccounts(conn)

	if err != nil {
		return err
	}

	validHomeId := "Такого Дома нет. "
	account.Valid = false

	for _, value := range houses {
		if value.Id != *account.HouseId {
			continue
		}

		validHomeId = ""
		account.Valid = true

		for _, value := range accounts {
			if value.Number == *account.Number {
				account.Text = "Такой лицевой счет уже есть. "
				account.Valid = false
			}
		}

		break
	}

	account.Text = strings.TrimSpace(validHomeId + account.Text)

	if !account.Valid {
		return errors.New(account.Text)
	}

	return nil
}

func ValidateDataAccountUpdate(conn *pgx.Conn, account dto.DataAccount) error {
	houses, err := repository.ReadHouses(conn)

	if err != nil {
		return err
	}

	accounts, err := repository.ReadAccounts(conn)

	if err != nil {
		return err
	}

	validId := "Такого ID нет. "
	account.Valid = false

	for _, value := range accounts {
		if value.Id != *account.Id {
			continue
		}

		validId = ""
		validHomeId := "Такого Дома нет. "

		for _, value := range houses {
			if value.Id != *account.HouseId {
				continue
			}

			validHomeId = ""
			account.Valid = true

			break
		}

		if value.Number == *account.Number &&
			value.HouseId == *account.HouseId &&
			(value.PremisesType == nil && account.PremisesType == nil ||
				value.PremisesType != nil && account.PremisesType != nil && *value.PremisesType == *account.PremisesType) &&
			(value.PremisesNumber == nil && account.PremisesNumber == nil ||
				value.PremisesNumber != nil && account.PremisesNumber != nil && *value.PremisesNumber == *account.PremisesNumber) {
			account.Text = account.Text + "Данные не изменились. "
			account.Valid = false
		}

		account.Text = validHomeId + account.Text

		break
	}

	account.Text = strings.TrimSpace(validId + account.Text)

	if !account.Valid {
		return errors.New(account.Text)
	}

	return nil
}

func ValidateDataAccountDelete(conn *pgx.Conn, account dto.DataAccount) error {
	accounts, err := repository.ReadAccounts(conn)

	if err != nil {
		return err
	}

	validId := "Такого ID нет. "
	account.Valid = false

	for _, value := range accounts {
		if value.Id != *account.Id {
			continue
		}

		validId = ""
		account.Valid = true

		break
	}

	account.Text = strings.TrimSpace(validId + account.Text)

	if !account.Valid {
		return errors.New(account.Text)
	}

	return nil
}

func ValidateDataDebtCreate(conn *pgx.Conn, debt dto.DataDebt) error {
	accounts, err := repository.ReadAccounts(conn)

	if err != nil {
		return err
	}

	validAccountNumber := "Такого лицевого счета нет. "
	debt.Valid = false

	for _, value := range accounts {
		if value.Number != *debt.AccountNumber {
			continue
		}

		validAccountNumber = ""
		debt.Valid = true

		break
	}

	debt.Text = strings.TrimSpace(validAccountNumber + debt.Text)

	if !debt.Valid {
		return errors.New(debt.Text)
	}

	return nil
}

func ValidateDataDebtUpdate(conn *pgx.Conn, debt dto.DataDebt) error {
	accounts, err := repository.ReadAccounts(conn)

	if err != nil {
		return err
	}

	debts, err := repository.ReadDebts(conn)

	if err != nil {
		return err
	}

	validId := "Такого ID нет. "
	debt.Valid = false

	for _, value := range debts {
		if value.Id != *debt.Id {
			continue
		}

		validId = ""
		validAccountNumber := "Такого Лицевого счета нет. "

		for _, value := range accounts {
			if value.Number != *debt.AccountNumber {
				continue
			}

			validAccountNumber = ""
			debt.Valid = true

			break
		}

		if value.AccountNumber == *debt.AccountNumber && value.OpeningBalance == *debt.OpeningBalance && value.Accrued == *debt.Accrued && value.Paid == *debt.Paid && value.ClosingBalance == *debt.ClosingBalance {
			debt.Text = debt.Text + "Данные не изменились. "
			debt.Valid = false
		}

		debt.Text = validAccountNumber + debt.Text

		break
	}

	debt.Text = strings.TrimSpace(validId + debt.Text)

	if !debt.Valid {
		return errors.New(debt.Text)
	}

	return nil
}

func ValidateDataDebtDelete(conn *pgx.Conn, debt dto.DataDebt) error {
	debts, err := repository.ReadDebts(conn)

	if err != nil {
		return err
	}

	validId := "Такого ID нет. "
	debt.Valid = false

	for _, value := range debts {
		if value.Id != *debt.Id {
			continue
		}

		validId = ""
		debt.Valid = true

		break
	}

	debt.Text = strings.TrimSpace(validId + debt.Text)

	if !debt.Valid {
		return errors.New(debt.Text)
	}

	return nil
}
