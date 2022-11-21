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

func (i *placeUseCase) AddPlace(place *entities.PlacePost) error {
	return i.storage.AddPlace(place)
}

func (i *placeUseCase) GetAllPlaces() (places []entities.Place, err error) {
	return i.storage.GetAllPlaces()
}

func (i *placeUseCase) GetPlace(id int) (place entities.Place, err error) {
	return i.storage.GetPlace(id)
}

func (i *placeUseCase) UpdatePlace(place *entities.PlacePost, id int) error {
	return i.storage.UpdatePlace(place, id)
}

func (i *placeUseCase) DeletePlace(id int) error {
	return i.storage.DeletePlace(id)
}
