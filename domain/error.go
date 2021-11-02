package domain

import "errors"

var (
	ErrNotFound error = errors.New("the requested resource doesn't exist")
)
