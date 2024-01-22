package domain

import "github.com/google/uuid"

type Bill struct {
	id uuid.UUID
	name string
	balance float64
	isClosed bool
}