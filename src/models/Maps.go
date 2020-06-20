package models

// Maps defines a basic structure for a map
type Maps struct {
	LatitudeMax  float32     `json:"MapLatitudeMax"`
	LatitudeMin  float32     `json:"MapLatitudeMin"`
	LongitudeMax float32     `json:"MapLongitudeMax"`
	LongitudeMin float32     `json:"MapLongitudeMin"`
	Lengthblocks MapDivision `json:"MapDivision"`
}

// MapDivision represents the length (latitude X longitude) of a given area in the map
type MapDivision struct {
	LengthLatitude  float32 `json:"LengthLatitude"`
	LengthLongitude float32 `json:"LengthLongitude"`
}
