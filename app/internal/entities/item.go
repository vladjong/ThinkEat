package entities

import "github.com/shopspring/decimal"

type Item struct {
	ID       int             `json:"id" db:"id,omitempty"`
	Name     string          `json:"name" db:"name,omitempty"`
	Describe string          `json:"describe" db:"describe"`
	Price    decimal.Decimal `json:"price" db:"price"`
	Weight   float64         `json:"weight" db:"weight"`
	Photo    string          `json:"photo" db:"photo"`
	Category Category        `json:"category"`
	PlaceId  int             `json:"place_id" db:"place_id"`
}

type Category struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name,omitempty"`
}
