package models

import "time"

type Clients struct {
	ID        int       `json:"ID" schema: "ID"`
	Name      string    `json:"Name" schema:"Name"`
	CreatedAt time.Time `schema:"-"`
}
