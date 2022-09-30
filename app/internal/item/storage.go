package item

import (
	"context"
)

type Storage interface {
	Create(ctx context.Context, item Item) (string, error)
	FindOnly(ctx context.Context, id string) (Item, error)
	Update(ctx context.Context, item Item) error
	Delete(ctx context.Context, id string) error
}
