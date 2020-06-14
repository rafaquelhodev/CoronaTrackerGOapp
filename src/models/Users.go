package models

import (
	"time"
)

// Users model for user
type Users struct {
	ID        int       `schema: "-"`
	Body      string    `schema:"body"`
	CreatedAt time.Time `schema:"-"`
	UpdatedAt time.Time `schema:"-"`
}
