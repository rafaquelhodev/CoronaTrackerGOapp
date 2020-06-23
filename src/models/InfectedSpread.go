package models

import "time"

// InfectedSpreadPeriod struct of an infected client
// InfectionPeriod = [initial spread date, final spread date]
type InfectedSpreadPeriod struct {
	IDclient        int
	InfectionPeriod [2]time.Time
}

// InfectedSpreadPeriodSlice is a slice of InfectedSpreadPeriod struct
type InfectedSpreadPeriodSlice []InfectedSpreadPeriod

// Implementing method to sort an array of InfectedSpreadPeriod according to infection date:
func (p InfectedSpreadPeriodSlice) Len() int {
	return len(p)
}

func (p InfectedSpreadPeriodSlice) Less(i, j int) bool {
	return (p[i].InfectionPeriod)[0].Before((p[j].InfectionPeriod)[0])
}

func (p InfectedSpreadPeriodSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
