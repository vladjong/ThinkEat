package postgressql

import "errors"

var (
	ErrItemsEmpty   = errors.New("items empty")
	ErrItemNotFound = errors.New("item not found")
)
