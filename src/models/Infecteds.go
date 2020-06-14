package models

// Infecteds infected person
type Infecteds struct {
	ID          int    `json:"ID" schema: "ID"`
	TestingDate string `json:"TestingDate" schema: "TestingDate"`
}
