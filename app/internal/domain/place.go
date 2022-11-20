package domain

import (
	"github.com/vladjong/ThinkEat/internal/adapters/db"
	"github.com/vladjong/ThinkEat/internal/entities"
)

type placeUseCase struct {
	storage db.PlaceStorager
}

func NewPlace(storage db.PlaceStorager) *placeUseCase {
	return &placeUseCase{storage: storage}
}

func (i *placeUseCase) AddPlace(place *entities.Place) error {
	return i.storage.AddPlace(place)
}
