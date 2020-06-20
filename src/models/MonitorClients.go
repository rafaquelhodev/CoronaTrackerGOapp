package models

import "time"

// MonitorClients data of a client
type MonitorClients struct {
	ID            int       `json:"ID" schema: "ID"`
	IDclient      int       `json:"IDclient" schema: "IDclient"`
	Time          time.Time `json:"Time" schema:"Time"`
	Latitude      float32   `json:"Xcoord" schema:"Latitude"`
	Longitude     float32   `json:"Ycoord" schema:"Longitude"`
	LocationBlock int       `json:"LocationBlock" schema:"LocationBlock"`
	Name          string    `json:"Name" schema:"Name"`
}

// GetAreaNumberInMap gets the location of the client in Earth
// The current position (latitude and longitude) is transformed in a matrix positioning
// Index i and j refer to longitude and latitude, respectively
func (client *MonitorClients) GetAreaNumberInMap(clientPosition Position, mapEarth Maps) {
	lengthLongitude := mapEarth.Lengthblocks.LengthLongitude
	LengthLatitude := mapEarth.Lengthblocks.LengthLatitude

	imax := int((mapEarth.LongitudeMax - mapEarth.LongitudeMin) / lengthLongitude)
	jmax := int((mapEarth.LatitudeMax - mapEarth.LatitudeMin) / LengthLatitude)

	i := int((clientPosition.Longitude-mapEarth.LongitudeMin)/lengthLongitude + 1)
	if i > imax {
		i = imax
	}

	j := int((clientPosition.Latitude-mapEarth.LatitudeMin)/LengthLatitude + 1)
	if j > jmax {
		j = jmax
	}

	(*client).LocationBlock = (i-1)*jmax + j
}
