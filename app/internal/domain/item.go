package domain

import (
	"github.com/vladjong/ThinkEat/internal/adapters/db"
	"github.com/vladjong/ThinkEat/internal/controller/http/handler/dto"
	"github.com/vladjong/ThinkEat/internal/entities"
)

type itemUseCase struct {
	storage db.ItemStorager
}

func NewItem(storage db.ItemStorager) *itemUseCase {
	return &itemUseCase{storage: storage}
}

func (i *itemUseCase) AddItem(itemDto dto.ItemDto) error {
	item := entities.Item{
		Name:     itemDto.Name,
		Describe: itemDto.Describe,
		Price:    itemDto.Price,
		Weight:   itemDto.Weight,
		Photo:    itemDto.Photo,
		Type:     itemDto.Type,
		PlaceId:  itemDto.PlaceId,
	}
	return i.storage.AddItem(item)
}

func (i *itemUseCase) GetAllItems() (items []entities.Item, err error) {
	return i.storage.GetAllItems()
}

func (i *itemUseCase) GetItem(id int) (item entities.Item, err error) {
	return i.storage.GetItem(id)
}

func (i *itemUseCase) UpdateItem(itemDto dto.ItemDto, id int) error {
	item := entities.Item{
		Name:     itemDto.Name,
		Describe: itemDto.Describe,
		Price:    itemDto.Price,
		Weight:   itemDto.Weight,
		Photo:    itemDto.Photo,
		Type:     itemDto.Type,
		PlaceId:  itemDto.PlaceId,
	}
	return i.storage.UpdateItem(item, id)
}

func (i *itemUseCase) DeleteItem(id int) error {
	return i.storage.DeleteItem(id)
}
