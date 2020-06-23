package models

// VirusSpread it is a struct that contains information about a virus
type VirusSpread struct {
	InfectionPeriod       int     `json="Infection_Period"`
	ContactTimeMinutes    float64 `json="Contact_Time_Minutes"`
	ContactDistanceMeters float64 `json="Contact_Distance_Meters"`
}
