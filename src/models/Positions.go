package models

// Position position of a client in terms of latitude and longitude
type Position struct {
	UserID    int     `json:"user"`
	Latitude  float32 `json:"lati"`
	Longitude float32 `json:"long"`
}
