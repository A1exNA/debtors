package dto

type Account struct {
	AccountDelete
	AccountCreate
}

type AccountCreate struct {
	Number         int     `json:"number"`
	HouseId        int     `json:"houseId"`
	PremisesType   *string `json:"premisesType"`
	PremisesNumber *string `json:"premisesNumber"`
}

type AccountDelete struct {
	Id int `json:"id"`
}
