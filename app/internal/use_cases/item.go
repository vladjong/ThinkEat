package usecases

import (
	"context"

	"github.com/vladjong/ThinkEat/internal/entities"
	"github.com/vladjong/ThinkEat/internal/interfaces"
)

type itemUseCase struct {
	storage interfaces.ItemStorages
}

func NewItemUseCase(storage interfaces.ItemStorages) *itemUseCase {
	return &itemUseCase{storage: storage}
}

func (s *itemUseCase) Create(ctx context.Context, item *entities.Item) (string, error) {
	return s.storage.Create(ctx, item)
}

func (s *itemUseCase) GetAll(ctx context.Context) ([]*entities.Item, error) {
	return s.storage.GetAll(ctx)
}

func (s *itemUseCase) GetName(ctx context.Context, name string) ([]*entities.Item, error) {
	return s.storage.GetName(ctx, name)
}

func (s *itemUseCase) GetID(ctx context.Context, id string) (*entities.Item, error) {
	return s.storage.GetID(ctx, id)
}

func (s *itemUseCase) Update(ctx context.Context, item *entities.Item) error {
	return s.storage.Update(ctx, item)
}

func (s *itemUseCase) Delete(ctx context.Context, id string) error {
	return s.storage.Delete(ctx, id)
}
