package models

import "testing"

func TestGetAreaNumberInMap(t *testing.T) {

	mapDivision := MapDivision{
		LengthLatitude:  10.0,
		LengthLongitude: 10.0,
	}

	mapEarth := Maps{
		LatitudeMax:  40.0,
		LatitudeMin:  -40.0,
		LongitudeMax: 20.0,
		LongitudeMin: -20.0,
		Lengthblocks: mapDivision,
	}

	tt := []struct {
		name           string
		clientPosition Position
		clientLocation int
	}{
		{
			name: "Negative latitude and longitude",
			clientPosition: Position{
				Longitude: -15.0,
				Latitude:  -35.0},
			clientLocation: 1,
		},
		{
			name: "Negative latitude and longitude. Longitude at a border",
			clientPosition: Position{
				Longitude: -10.0,
				Latitude:  -35.0},
			clientLocation: 9,
		},
		{
			name: "Negative latitude and longitude. Longitude at extreme left border",
			clientPosition: Position{
				Longitude: -20.0,
				Latitude:  -35.0},
			clientLocation: 1,
		},
		{
			name: "Negative latitude and longitude. Latitude at a border",
			clientPosition: Position{
				Longitude: -15.0,
				Latitude:  -20.0},
			clientLocation: 3,
		},
		{
			name: "Positive longitude and negative latitude",
			clientPosition: Position{
				Longitude: 15.0,
				Latitude:  -35.0},
			clientLocation: 25,
		},
		{
			name: "Positive longitude and negative latitude. Longitude at extreme right border",
			clientPosition: Position{
				Longitude: 20.0,
				Latitude:  -35.0},
			clientLocation: 25,
		},
		{
			name: "Positive longitude and negative latitude. Latitude at a border",
			clientPosition: Position{
				Longitude: 15.0,
				Latitude:  -30.0},
			clientLocation: 26,
		},
		{
			name: "Positive latitude and negative longitude",
			clientPosition: Position{
				Longitude: -15.0,
				Latitude:  35.0},
			clientLocation: 8,
		},
		{
			name: "Positive latitude and negative longitude. Latitude extreme upper",
			clientPosition: Position{
				Longitude: -15.0,
				Latitude:  40.0},
			clientLocation: 8,
		},
		{
			name: "Positive latitude and negative longitude. Latitude extreme lower",
			clientPosition: Position{
				Longitude: -15.0,
				Latitude:  -40.0},
			clientLocation: 1,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			monitoredClient := &MonitorClients{}

			monitoredClient.GetAreaNumberInMap(tc.clientPosition, mapEarth)

			if monitoredClient.LocationBlock != tc.clientLocation {
				t.Errorf("Expected client location = %d; got %d", tc.clientLocation, monitoredClient.LocationBlock)
			}
		})
	}
}
