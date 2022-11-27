package dto

import "github.com/shopspring/decimal"

type ItemDto struct {
	ID       int             `json:"id"`
	Name     string          `json:"name"`
	Describe string          `json:"describe"`
	Price    decimal.Decimal `json:"price"`
	Weight   float64         `json:"weight"`
	Photo    string          `json:"photo"`
	Type     string          `json:"type"`
	PlaceId  int             `json:"place_id"`
}
