package models

type Weather struct {
	Status Status `json:"status"`
}

type Status struct {
	Water uint8 `json:"water"`
	Wind  uint8 `json:"wind"`
}