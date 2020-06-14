package models

// Clients data of a person
type Clients struct {
	ID       int     `json:"ID" schema: "ID"`
	IDclient int     `json:"IDclient" schema: "IDclient"`
	Time     string  `json:"Time" schema:"Time"`
	Xcoord   float64 `json:"Xcoord" schema:"Xcoord"`
	Ycoord   float64 `json:"Ycoord" schema:"Ycoord"`
	City     string  `json:"City" schema:"City"`
	Day      string  `json:"Day" schema:"Day"`
	Name     string  `json:"Name" schema:"Name"`
}
