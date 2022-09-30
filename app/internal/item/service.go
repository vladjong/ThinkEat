package item

import (
	"context"

	"github.com/vladjong/ThinkEat/pkg/logging"
)

type Service struct {
	storage Storage
	logger  *logging.Logger
}

func (s *Service) Create(ctx context.Context, dto ItemDTO) (Item, error) {
	return Item{}, nil
}
