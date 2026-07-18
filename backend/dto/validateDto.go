package dto

import "time"

type MessageError struct {
	Message string
}

type DataHouse struct {
	ValueHouse
	Valid bool
	Text  string
}

type ValueHouse struct {
	Id         *int
	Address    *string
	IsServiced *bool
}

type DataAccount struct {
	ValueAccount
	Valid bool
	Text  string
}

type ValueAccount struct {
	Id             *int
	Number         *int
	HouseId        *int
	PremisesType   *string
	PremisesNumber *string
	OwnerName      *string
	OwnerPhone     *string
}

type DataDebt struct {
	ValueDebt
	Valid bool
	Text  string
}

type ValueDebt struct {
	Id             *int
	AccountNumber  *int
	ReportDate     *time.Time
	OpeningBalance *float64
	Accrued        *float64
	Paid           *float64
	ClosingBalance *float64
	UploadDate     *time.Time
}
