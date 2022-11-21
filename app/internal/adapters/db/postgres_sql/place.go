package postgressql

import (
	"errors"

	"github.com/vladjong/ThinkEat/internal/entities"
)

func (s *thinkEatStorage) AddPlace(place *entities.PlacePost) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	var idContact int
	queryContact := `INSERT INTO contacts (city, street)
				VALUES ($1, $2) RETURNING id;`
	row := tx.QueryRow(queryContact, place.Contact.City, place.Contact.Street)
	if err := row.Scan(&idContact); err != nil {
		if rb := tx.Rollback(); rb != nil {
			return rb
		}
		return err
	}
	for _, value := range place.Contact.WorkSchedule {
		querySchedule := `INSERT INTO work_schedule (weekday, beggin_time, end_time, is_day_off, contact_id)
							VALUES ($1, $2, $3, $4, $5);`
		if _, err := tx.Exec(querySchedule, value.Weekday, value.BegginTime,
			value.EndTime, value.IsDayOff, idContact); err != nil {
			if rb := tx.Rollback(); rb != nil {
				return rb
			}
			return err
		}
	}
	queryItem := `INSERT INTO places (name, describe, url, photo, class, contact_id)
					VALUES ($1, $2, $3, $4, $5, $6);`
	if _, err := tx.Exec(queryItem, place.Name, place.Describe, place.URL, place.Photo, place.Class, idContact); err != nil {
		if rb := tx.Rollback(); rb != nil {
			return rb
		}
		return err
	}
	return tx.Commit()
}

func (s *thinkEatStorage) GetAllPlaces() (places []entities.Place, err error) {
	query := `SELECT p.id, name, describe, url, photo, class, city, street
				FROM places AS p
				JOIN contacts c ON p.contact_id = c.id;`
	if err := s.db.Select(&places, query); err != nil {
		return nil, err
	}
	if len(places) == 0 {
		return nil, errors.New("places don't exist")
	}
	return places, nil
}

func (s *thinkEatStorage) GetPlace(id int) (place entities.Place, err error) {
	query := `SELECT p.id, name, describe, url, photo, class, city, street
				FROM places AS p
				JOIN contacts c ON p.contact_id = c.id
				WHERE p.id = $1;`
	if err := s.db.Get(&place, query, id); err != nil {
		return place, err
	}
	return place, nil
}

func (s *thinkEatStorage) UpdatePlace(place *entities.PlacePost, id int) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	queryItem := `UPDATE places
					SET name = $1, describe = $2, url = $3, photo = $4, class = $5
					WHERE id = $6`
	if _, err := tx.Exec(queryItem, place.Name, place.Describe, place.URL, place.Photo, place.Class, id); err != nil {
		if rb := tx.Rollback(); rb != nil {
			return rb
		}
		return err
	}
	return tx.Commit()
}

func (s *thinkEatStorage) DeletePlace(id int) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	var idPlace int
	queruSelect := `SELECT id FROM places WHERE id = $1`
	if err := s.db.Get(&idPlace, queruSelect, id); err != nil {
		return err
	}
	query := `DELETE FROM places WHERE id = $1`
	if _, err = tx.Exec(query, id); err != nil {
		if rb := tx.Rollback(); rb != nil {
			return rb
		}
		return err
	}
	return tx.Commit()
}
