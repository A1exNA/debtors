package handler

import (
	"strings"
	"time"
	"unicode/utf8"
)

type Value struct {
	Id             *int       // Общая
	Address        *string    // Houses
	AccountNumber  *string    // Accounts
	HouseId        *int       //
	PremisesType   *string    //
	PremisesNumber *string    //
	OwnerName      *string    //
	OwnerPhone     *string    //
	ReportDate     *time.Time // Debts
}

type Data struct {
	Value
	Valid bool
	Text  string
}

func ValidateData(data Data) Data {
	if data.Id != nil {
		if *data.Id <= 0 {
			data.Valid = false
			data.Text = data.Text + "ID не может быть меньше или ровно нулю. "
		}
	}

	if data.Address != nil {
		if *data.Address == "" {
			data.Valid = false
			data.Text = data.Text + "Address не введен. "
		} else if utf8.RuneCountInString(*data.Address) > 50 {
			data.Valid = false
			data.Text = data.Text + "Address слишком длинный. "
		}
	}

	if data.AccountNumber != nil {
		if *data.AccountNumber == "" {
			data.Valid = false
			data.Text = data.Text + "AccountNumber не введен. "
		} else if utf8.RuneCountInString(*data.AccountNumber) > 10 {
			data.Valid = false
			data.Text = data.Text + "AccountNumber слишком длинный. "
		}
	}

	if data.HouseId != nil {
		if *data.HouseId <= 0 {
			data.Valid = false
			data.Text = data.Text + "HouseID не может быть меньше нуля. "
		}
	}

	if data.PremisesType != nil {
		if *data.PremisesType == "" {
			data.Valid = false
			data.Text = data.Text + "PremisesType не введен. "
		} else if utf8.RuneCountInString(*data.PremisesType) > 15 {
			data.Valid = false
			data.Text = data.Text + "PremisesType слишком длинный. "
		}
	}

	if data.PremisesNumber != nil {
		if *data.PremisesNumber == "" {
			data.Valid = false
			data.Text = data.Text + "PremisesNumber не введен. "
		} else if utf8.RuneCountInString(*data.PremisesNumber) > 10 {
			data.Valid = false
			data.Text = data.Text + "PremisesNumber слишком длинный. "
		}
	}

	if data.OwnerName != nil {
		if *data.OwnerName == "" {
			data.Valid = false
			data.Text = data.Text + "OwnerName не введен. "
		} else if utf8.RuneCountInString(*data.OwnerName) > 100 {
			data.Valid = false
			data.Text = data.Text + "OwnerName слишком длинный. "
		}
	}

	if data.OwnerPhone != nil {
		if *data.OwnerPhone == "" {
			data.Valid = false
			data.Text = data.Text + "OwnerPhone не введен. "
		} else if utf8.RuneCountInString(*data.OwnerPhone) > 15 {
			data.Valid = false
			data.Text = data.Text + "OwnerPhone слишком длинный. "
		}
	}

	if data.ReportDate != nil {
		if data.ReportDate.IsZero() {
			data.Valid = false
			data.Text = data.Text + "Некорректная дата. "
		}
	}

	data.Text = strings.TrimSpace(data.Text)

	return data
}
