package domain

import "github.com/vladjong/ThinkEat/internal/entities"

type ItemDomain interface {
	AddItem(item *entities.Item) error
}
