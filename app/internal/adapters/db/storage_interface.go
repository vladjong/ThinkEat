package db

import "github.com/vladjong/ThinkEat/internal/entities"

type ItemStorager interface {
	AddItem(item *entities.Item) error
}

type PlaceStorager interface {
	AddPlace(place *entities.Place) error
}
