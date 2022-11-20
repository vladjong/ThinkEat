package postgressql

import (
	"github.com/vladjong/ThinkEat/internal/entities"
)

func (s *thinkEatStorage) AddItem(item *entities.Item) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	var id int
	queryCategory := `INSERT INTO categories (name)
						VALUES ($1) ON CONFLICT (name)
						DO NOTHING RETURNING id`
	row := tx.QueryRow(queryCategory, item.Category.Name)
	if err := row.Scan(&id); err != nil {
		if rb := tx.Rollback(); rb != nil {
			return rb
		}
		return err
	}
	queryItem := `INSERT INTO items (name, describe, price, weight, photo, category_id, place_id)
				VALUES ($1, $2, $3, $4, $5, $6, $7)`
	if _, err := tx.Exec(queryItem, item.Name, item.Describe, item.Price, item.Weight, item.Photo, 3, item.PlaceId); err != nil {
		if rb := tx.Rollback(); rb != nil {
			return rb
		}
		return err
	}
	return tx.Commit()
}
