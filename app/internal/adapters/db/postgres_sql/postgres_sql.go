package postgressql

import "github.com/jmoiron/sqlx"

type thinkEatStorage struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *thinkEatStorage {
	return &thinkEatStorage{
		db: db,
	}
}
