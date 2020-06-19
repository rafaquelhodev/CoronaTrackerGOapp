package models

import (
	"encoding/json"
	"time"
)

// Infecteds infected person
type Infecteds struct {
	ID          int       `json:"ID" schema: "ID"`
	IDclient    int       `json:"IDclient" schema: "IDclient"`
	TestingDate time.Time `json:"TestingDate" schema: "TestingDate"`
}

func (n *Infecteds) UnmarshalJSON(bytes []byte) error {
	var timestr string
	timeparsed, _ := time.Parse("2006-01-02", timestr)
	err := json.Unmarshal(bytes, &timeparsed)
	if err != nil {
		return err
	}
	n.TestingDate = timeparsed
	return nil
}
