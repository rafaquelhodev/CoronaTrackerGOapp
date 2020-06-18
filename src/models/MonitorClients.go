package models

import "time"

// Clients data of a person
type MonitorClients struct {
	ID        int       `json:"ID" schema: "ID"`
	IDclient  int       `json:"IDclient" schema: "IDclient"`
	Time      time.Time `json:"Time" schema:"Time"`
	Latitude  float32   `json:"Xcoord" schema:"Latitude"`
	Longitude float32   `json:"Ycoord" schema:"Longitude"`
	City      string    `json:"City" schema:"City"`
	Day       string    `json:"Day" schema:"Day"`
	Name      string    `json:"Name" schema:"Name"`
}
