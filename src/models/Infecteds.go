package models

import (
	"time"
)

// Infecteds infected person
type Infecteds struct {
	ID          int       `json:"ID" schema: "ID"`
	IDclient    int       `json:"IDclient" schema: "IDclient"`
	TestingDate time.Time `json:"TestingDate" schema: "TestingDate"`
}
