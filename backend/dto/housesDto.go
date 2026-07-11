package dto

type House struct {
	Id         int
	Address    string
	IsServiced bool
}

type HouseCreate struct {
	Address    string
	IsServiced bool
}

type HouseDelete struct {
	Id int
}
