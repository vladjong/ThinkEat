package domain

import (
	"github.com/vladjong/ThinkEat/internal/controller/http/handler/dto"
	"github.com/vladjong/ThinkEat/internal/entities"
)

type ItemDomain interface {
	AddItem(item dto.ItemDto) error
	GetAllItems() (items []entities.Item, err error)
	GetItem(id int) (item entities.Item, err error)
	UpdateItem(itemDto dto.ItemDto, id int) error
	DeleteItem(id int) error
}

type PlaceDomain interface {
	AddPlace(place *entities.PlacePost) error
	GetAllPlaces() (places []entities.Place, err error)
	GetPlace(id int) (place entities.Place, err error)
	UpdatePlace(place *entities.PlacePost, id int) error
	DeletePlace(id int) error
}
