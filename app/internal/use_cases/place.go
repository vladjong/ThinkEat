package usecase

// import (
// 	"context"

// 	"github.com/vladjong/ThinkEat/internal/entities"
// 	"github.com/vladjong/ThinkEat/internal/interfaces"
// )

// type placeUseCase struct {
// 	storage interfaces.PlaceStorages
// }

// func NewPlaceUseCase(storage interfaces.PlaceStorages) *placeUseCase {
// 	return &placeUseCase{storage: storage}
// }

// func (p *placeUseCase) Create(ctx context.Context, place *entities.Place) (string, error) {
// 	return p.storage.Create(ctx, place)
// }

// func (p *placeUseCase) GetAll(ctx context.Context) ([]*entities.Place, error) {
// 	return p.storage.GetAll(ctx)
// }

// func (p *placeUseCase) GetName(ctx context.Context, name string) ([]*entities.Place, error) {
// 	return p.storage.GetName(ctx, name)
// }

// func (p *placeUseCase) GetID(ctx context.Context, id string) (*entities.Place, error) {
// 	return p.storage.GetID(ctx, id)
// }

// func (p *placeUseCase) Update(ctx context.Context, place *entities.Place) error {
// 	return p.storage.Update(ctx, place)
// }

// func (p *placeUseCase) Delete(ctx context.Context, id string) error {
// 	return p.storage.Delete(ctx, id)
// }
