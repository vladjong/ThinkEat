package service

import (
	"context"

	"github.com/vladjong/ThinkEat/internal/domain/entity"
)

type ItemStorage interface {
	Create(ctx context.Context, item *entity.Item) (string, error)
	GetAll(ctx context.Context) ([]*entity.Item, error)
	GetName(ctx context.Context, name string) ([]*entity.Item, error)
	GetID(ctx context.Context, id string) (*entity.Item, error)
	Update(ctx context.Context, item *entity.Item) error
	Delete(ctx context.Context, id string) error
}

type itemService struct {
	storage ItemStorage
}

func NewItemService(storage ItemStorage) *itemService {
	return &itemService{storage: storage}
}

func (s *itemService) Create(ctx context.Context) entity.Item {
	return entity.Item{}
}

func (s *itemService) GetAll(ctx context.Context) ([]*entity.Item, error) {
	return s.storage.GetAll(ctx)
}

func (s *itemService) GetName(ctx context.Context, name string) ([]*entity.Item, error) {
	return s.storage.GetName(ctx, name)
}

func (s *itemService) GetID(ctx context.Context, id string) (*entity.Item, error) {
	return s.storage.GetID(ctx, id)
}

func (s *itemService) Update(ctx context.Context, item *entity.Item) error {
	return s.storage.Update(ctx, item)
}

func (s *itemService) Delete(ctx context.Context, id string) error {
	return s.storage.Delete(ctx, id)
}
