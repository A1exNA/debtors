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
	OwnerName      *string `json:"ownerName"`
	OwnerPhone     *string `json:"ownerPhone"`
}

type AccountDelete struct {
	Id int `json:"id"`
}
