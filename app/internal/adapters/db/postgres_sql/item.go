package postgressql

import (
	"github.com/vladjong/ThinkEat/internal/entities"
)

func (s *thinkEatStorage) AddItem(item entities.Item) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	var idItem int
	queryItem := `INSERT INTO items (name, describe, price, weight, photo, type, place_id)
				VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;`
	row := tx.QueryRow(queryItem, item.Name, item.Describe, item.Price, item.Weight, item.Photo, item.Type, item.PlaceId)
	if err := row.Scan(&idItem); err != nil {
		if rb := tx.Rollback(); rb != nil {
			return rb
		}
		return err
	}
	return tx.Commit()
}

func (s *thinkEatStorage) GetAllItems() (items []entities.Item, err error) {
	queryItems := `SELECT * FROM items;`
	if err := s.db.Select(&items, queryItems); err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, ErrItemsEmpty
	}
	return items, nil
}

func (s *thinkEatStorage) GetItem(id int) (item entities.Item, err error) {
	queryItems := `SELECT * FROM items WHERE id = $1;`
	if err := s.db.Get(&item, queryItems, id); err != nil {
		return item, ErrItemNotFound
	}
	return item, nil
}

func (s *thinkEatStorage) UpdateItem(item entities.Item, id int) error {
	queryItem := `UPDATE items
					SET name = $1, describe = $2, price = $3, weight = $4, photo = $5, type = $6, place_id = $7
					WHERE id = $8;`
	if _, err := s.db.Exec(queryItem, item.Name, item.Describe, item.Price, item.Weight, item.Photo, item.Type, item.PlaceId, id); err != nil {
		return err
	}
	return nil
}

func (s *thinkEatStorage) DeleteItem(id int) error {
	if _, err := s.GetItem(id); err != nil {
		return err
	}
	query := `DELETE FROM items WHERE id = $1`
	if _, err := s.db.Exec(query, id); err != nil {
		return err
	}
	return nil
}
