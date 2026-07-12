package dto

import "time"

type Debt struct {
	DebtDelete
	DebtCreate
}

type DebtCreate struct {
	AccountNumber  string    `json:"accountNumber"`
	ReportDate     time.Time `json:"reportDate"`
	OpeningBalance float64   `json:"openingBalance"`
	Accrued        float64   `json:"accrued"`
	Paid           float64   `json:"paid"`
	ClosingBalance float64   `json:"closingBalance"`
}

type DebtDelete struct {
	Id int `json:"id"`
}
