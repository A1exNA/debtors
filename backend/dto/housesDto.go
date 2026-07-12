package dto

type House struct {
	HouseDelete
	HouseCreate
}

type HouseCreate struct {
	Address    string `json:"address"`
	IsServiced bool   `json:"isServiced"`
}

type HouseDelete struct {
	Id int `json:"id"`
}
