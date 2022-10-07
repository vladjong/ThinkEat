package interfaces

import (
	"context"

	"github.com/vladjong/ThinkEat/internal/entities"
)

type (
	ItemStorages interface {
		Create(ctx context.Context, item *entities.Item) (string, error)
		GetAll(ctx context.Context) ([]*entities.Item, error)
		GetName(ctx context.Context, name string) ([]*entities.Item, error)
		GetID(ctx context.Context, id string) (*entities.Item, error)
		Update(ctx context.Context, item *entities.Item) error
		Delete(ctx context.Context, id string) error
	}

	Item interface {
		Create(ctx context.Context, item *entities.Item) (string, error)
		GetAll(ctx context.Context) ([]*entities.Item, error)
		GetName(ctx context.Context, name string) ([]*entities.Item, error)
		GetID(ctx context.Context, id string) (*entities.Item, error)
		Update(ctx context.Context, item *entities.Item) error
		Delete(ctx context.Context, id string) error
	}
)
