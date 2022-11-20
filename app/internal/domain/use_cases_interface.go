package domain

import "github.com/vladjong/ThinkEat/internal/entities"

type ItemDomain interface {
	AddItem(item *entities.Item) error
}

type PlaceDomain interface {
	AddPlace(place *entities.Place) error
}
