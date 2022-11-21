package db

import "github.com/vladjong/ThinkEat/internal/entities"

type ItemStorager interface {
	AddItem(item *entities.Item) error
}

type PlaceStorager interface {
	AddPlace(place *entities.PlacePost) error
	GetAllPlaces() (places []entities.Place, err error)
	GetPlace(id int) (place entities.Place, err error)
	UpdatePlace(place *entities.PlacePost, id int) error
	DeletePlace(id int) error
}
