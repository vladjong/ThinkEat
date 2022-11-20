package postgressql

import (
	"github.com/vladjong/ThinkEat/internal/entities"
)

func (s *thinkEatStorage) AddPlace(place *entities.Place) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	var idSchedule int
	querySchedule := `INSERT INTO work_schedule (weekday, beggin_time, end_time, is_day_off)
						VALUES ($1, $2, $3, $4) RETURNING id`
	row := tx.QueryRow(querySchedule, place.Contacts.WorkSchedule.Weekend, place.Contacts.WorkSchedule.BegginTime,
		place.Contacts.WorkSchedule.EndTime, place.Contacts.WorkSchedule.IsDayOff)
	if err := row.Scan(&idSchedule); err != nil {
		if rb := tx.Rollback(); rb != nil {
			return rb
		}
		return err
	}
	var idContact int
	queryContact := `INSERT INTO contacts (city, street, work_schedule_id)
				VALUES ($1, $2, $3) RETURNING id`
	row = tx.QueryRow(queryContact, place.Contacts.City, place.Contacts.Street, idSchedule)
	if err := row.Scan(&idContact); err != nil {
		if rb := tx.Rollback(); rb != nil {
			return rb
		}
		return err
	}
	queryItem := `INSERT INTO places (name, describe, url, photo, class, contact_id)
					VALUES ($1, $2, $3, $4, $5, $6)`
	if _, err := tx.Exec(queryItem, place.Name, place.Describe, place.URL, place.Photo, place.Class, idContact); err != nil {
		if rb := tx.Rollback(); rb != nil {
			return rb
		}
		return err
	}
	return tx.Commit()
}
