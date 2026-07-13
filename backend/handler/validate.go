package handler

import (
	"strings"
	"time"
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
		}
	}

	if data.AccountNumber != nil {
		if *data.AccountNumber == "" {
			data.Valid = false
			data.Text = data.Text + "AccountNumber не введен. "
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
		}
	}

	if data.PremisesNumber != nil {
		if *data.PremisesNumber == "" {
			data.Valid = false
			data.Text = data.Text + "PremisesNumber не введен. "
		}
	}

	if data.OwnerName != nil {
		if *data.OwnerName == "" {
			data.Valid = false
			data.Text = data.Text + "OwnerName не введен. "
		}
	}

	if data.OwnerPhone != nil {
		if *data.OwnerPhone == "" {
			data.Valid = false
			data.Text = data.Text + "OwnerPhone не введен. "
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
