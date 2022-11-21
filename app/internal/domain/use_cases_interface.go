package domain

import "github.com/vladjong/ThinkEat/internal/entities"

type ItemDomain interface {
	AddItem(item *entities.Item) error
}

type PlaceDomain interface {
	AddPlace(place *entities.PlacePost) error
	GetAllPlaces() (places []entities.Place, err error)
	GetPlace(id int) (place entities.Place, err error)
	UpdatePlace(place *entities.PlacePost, id int) error
	DeletePlace(id int) error
}
