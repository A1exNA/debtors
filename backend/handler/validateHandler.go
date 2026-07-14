package handler

import (
	"debtors/dto"
	"errors"
	"strings"
	"unicode/utf8"
)

func ValidateDataHouse(house dto.DataHouse) error {
	if house.Id != nil {
		if *house.Id <= 0 {
			house.Valid = false
			house.Text = house.Text + "ID не может быть меньше или ровно нулю. "
		}
	}

	if house.Address != nil {
		if *house.Address == "" {
			house.Valid = false
			house.Text = house.Text + "Address не введен. "
		} else if utf8.RuneCountInString(*house.Address) > 50 {
			house.Valid = false
			house.Text = house.Text + "Address слишком длинный. "
		}
	}

	house.Text = strings.TrimSpace(house.Text)

	if house.Text != "" {
		return errors.New(house.Text)
	}

	return nil
}

func ValidateDataAccount(account dto.DataAccount) error {
	if account.Id != nil {
		if *account.Id <= 0 {
			account.Valid = false
			account.Text = account.Text + "ID не может быть меньше или ровно нулю. "
		}
	}

	if account.AccountNumber != nil {
		if *account.AccountNumber == "" {
			account.Valid = false
			account.Text = account.Text + "AccountNumber не введен. "
		} else if utf8.RuneCountInString(*account.AccountNumber) > 10 {
			account.Valid = false
			account.Text = account.Text + "AccountNumber слишком длинный. "
		}
	}

	if account.HouseId != nil {
		if *account.HouseId <= 0 {
			account.Valid = false
			account.Text = account.Text + "HouseID не может быть меньше нуля. "
		}
	}

	if account.PremisesType != nil {
		if *account.PremisesType == "" {
			account.Valid = false
			account.Text = account.Text + "PremisesType не введен. "
		} else if utf8.RuneCountInString(*account.PremisesType) > 15 {
			account.Valid = false
			account.Text = account.Text + "PremisesType слишком длинный. "
		}
	}

	if account.PremisesNumber != nil {
		if *account.PremisesNumber == "" {
			account.Valid = false
			account.Text = account.Text + "PremisesNumber не введен. "
		} else if utf8.RuneCountInString(*account.PremisesNumber) > 10 {
			account.Valid = false
			account.Text = account.Text + "PremisesNumber слишком длинный. "
		}
	}

	if account.OwnerName != nil {
		if *account.OwnerName == "" {
			account.Valid = false
			account.Text = account.Text + "OwnerName не введен. "
		} else if utf8.RuneCountInString(*account.OwnerName) > 100 {
			account.Valid = false
			account.Text = account.Text + "OwnerName слишком длинный. "
		}
	}

	if account.OwnerPhone != nil {
		if *account.OwnerPhone == "" {
			account.Valid = false
			account.Text = account.Text + "OwnerPhone не введен. "
		} else if utf8.RuneCountInString(*account.OwnerPhone) > 15 {
			account.Valid = false
			account.Text = account.Text + "OwnerPhone слишком длинный. "
		}
	}

	account.Text = strings.TrimSpace(account.Text)

	if account.Text != "" {
		return errors.New(account.Text)
	}

	return nil
}

func ValidateDataDebt(debt dto.DataDebt) error {
	if debt.Id != nil {
		if *debt.Id <= 0 {
			debt.Valid = false
			debt.Text = debt.Text + "ID не может быть меньше или ровно нулю. "
		}
	}

	if debt.AccountNumber != nil {
		if *debt.AccountNumber == "" {
			debt.Valid = false
			debt.Text = debt.Text + "AccountNumber не введен. "
		} else if utf8.RuneCountInString(*debt.AccountNumber) > 10 {
			debt.Valid = false
			debt.Text = debt.Text + "AccountNumber слишком длинный. "
		}
	}

	if debt.ReportDate != nil {
		if debt.ReportDate.IsZero() {
			debt.Valid = false
			debt.Text = debt.Text + "Некорректная дата. "
		}
	}

	if debt.OpeningBalance != nil {
		if *debt.OpeningBalance == 0 {
			debt.Valid = false
			debt.Text = debt.Text + "OpeningBalance не введен. "
		}
	}

	if debt.Accrued != nil {
		if *debt.Accrued == 0 {
			debt.Valid = false
			debt.Text = debt.Text + "OpeningBalance не введен. "
		}
	}

	if debt.Paid != nil {
		if *debt.Paid == 0 {
			debt.Valid = false
			debt.Text = debt.Text + "OpeningBalance не введен. "
		}
	}

	if debt.ClosingBalance != nil {
		if *debt.ClosingBalance == 0 {
			debt.Valid = false
			debt.Text = debt.Text + "ClosingBalance не введен. "
		}
	}

	debt.Text = strings.TrimSpace(debt.Text)

	if debt.Text != "" {
		return errors.New(debt.Text)
	}

	return nil
}
