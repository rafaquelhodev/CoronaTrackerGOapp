package controllers

import (
	"math"
	"models"
	"sort"
	"strings"
	"time"
)

// HandleOnInfecteds reads the infected table and calculates the initial and final spread dates for each infected
func (t *AdminController) HandleOnInfecteds(infectedPeople map[int]models.InfectedSpreadPeriod, virusSpread models.VirusSpread) {

	// Reading all data from infected table in DB
	infecteds := []*models.Infecteds{}
	t.Ctx.DB.Find(&infecteds)

	for _, personInfected := range infecteds {
		id := personInfected.IDclient

		testDate := personInfected.TestingDate

		virusSpreadPeriod := [2]time.Time{testDate.AddDate(0, 0, -virusSpread.InfectionPeriod),
			testDate.AddDate(0, 0, virusSpread.InfectionPeriod)}

		infectedPeople[id] = models.InfectedSpreadPeriod{
			IDclient:        id,
			InfectionPeriod: virusSpreadPeriod,
		}
	}
}

// RetrieveDataInfectedPerson retrieves data of an infected person and gets the map areas that might be infected
func (t *AdminController) RetrieveDataInfectedPerson(idclient int, initialSpreadDate time.Time,
	finalSpreadDate time.Time) (map[string][]*models.MonitorClients, map[int]bool) {

	dataInfected := make(map[string][]*models.MonitorClients, 0)
	mapAreasInfected := make(map[int]bool, 0)

	sqlQuery := `time BETWEEN ? AND ? AND idclient = ?`

	infectedSQL := []*models.MonitorClients{}
	t.Ctx.DB.Where(sqlQuery, initialSpreadDate, finalSpreadDate, idclient).Find(&infectedSQL)

	for _, infected := range infectedSQL {
		infectionDay := infected.Time.String()[0:10]

		dataInfected[infectionDay] = append(dataInfected[infectionDay], infected)
		mapAreasInfected[infected.LocationBlock] = true
	}

	return dataInfected, mapAreasInfected
}

// CommonAreaDate finds clients that were in the same map area and dates of the infected person
func (t *AdminController) CommonAreaDate(initialSpreadDate time.Time, finalSpreadDate time.Time,
	mapAreasInfected map[int]bool, infectedAnalysed map[int]models.InfectedSpreadPeriod) map[string][]*models.MonitorClients {

	dataClients := make(map[string][]*models.MonitorClients, 0)

	sqlQuery := `time BETWEEN ? AND ?
	AND idclient NOT IN (?` + strings.Repeat(",?", len(infectedAnalysed)-1) + `)
	AND location_block IN (?` + strings.Repeat(", ?", len(mapAreasInfected)-1) + `)`

	args := []interface{}{}

	dates := []interface{}{initialSpreadDate, finalSpreadDate}

	args = append(args, dates...)

	for idnow := range infectedAnalysed {
		args = append(args, idnow)
	}

	for cityNow := range mapAreasInfected {
		args = append(args, cityNow)
	}

	clients := []*models.MonitorClients{}
	t.Ctx.DB.Where(sqlQuery, args...).Find(&clients)

	for _, client := range clients {
		day := client.Time.String()[0:10]
		dataClients[day] = append(dataClients[day], client)
	}

	return dataClients
}

// FindOldestInfected finds the person that was infected first among the people in the map infectedPeople
func FindOldestInfected(infectedPeople map[int]models.InfectedSpreadPeriod) (int, time.Time, time.Time) {

	datesInfected := make(models.InfectedSpreadPeriodSlice, 0, len(infectedPeople))

	for _, d := range infectedPeople {
		datesInfected = append(datesInfected, d)
	}

	sort.Sort(datesInfected)

	idInfected := datesInfected[0].IDclient
	initialDateSpread := (datesInfected[0].InfectionPeriod)[0]
	finalDateSpread := (datesInfected[0].InfectionPeriod)[1]

	return idInfected, initialDateSpread, finalDateSpread
}

// distance calculates distance between two points using Haversin formula
func distance(lati1 float32, long1 float32, lati2 float32, long2 float32) float64 {
	var earthRadius float64
	earthRadius = 6371000.0 // meters

	var degToRad float32
	degToRad = math.Pi / 180.0

	omega1 := float64(lati1 * degToRad)
	omega2 := float64(lati2 * degToRad)
	deltaOmega := float64((lati2 - lati1) * degToRad)
	deltaLambda := float64((long2 - long1) * degToRad)

	a := math.Pow(math.Sin(deltaOmega*0.5), 2) +
		math.Cos(omega1)*math.Cos(omega2)*math.Pow(math.Sin(deltaLambda*0.5), 2)

	c := 2.0 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}

// contactWithInfected finds clients that were in contact with infected person
func contactWithInfected(infectedData map[string][]*models.MonitorClients, possibleInfecteds map[string][]*models.MonitorClients,
	virusSpread models.VirusSpread) map[int][]*models.MonitorClients {

	clientsContactInfected := make(map[int][]*models.MonitorClients, 0)

	for day, infectedClient := range infectedData {

		clientsPossibleInfecteds, foundPossibleClient := possibleInfecteds[day]

		if !foundPossibleClient {
			continue
		}

		for indexInfected := range infectedClient {
			latInfected := infectedClient[indexInfected].Latitude
			longInfected := infectedClient[indexInfected].Longitude
			timeInfected := infectedClient[indexInfected].Time

			for _, clientPossibleInfected := range clientsPossibleInfecteds {
				latPossible := clientPossibleInfected.Latitude
				longPossible := clientPossibleInfected.Longitude
				timePossible := clientPossibleInfected.Time

				dist := distance(latInfected, longInfected, latPossible, longPossible)

				timeDiff := timeInfected.Sub(timePossible).Minutes()

				if dist <= virusSpread.ContactDistanceMeters && math.Abs(timeDiff) <= virusSpread.ContactTimeMinutes {
					clientsContactInfected[clientPossibleInfected.IDclient] = append(clientsContactInfected[clientPossibleInfected.IDclient], clientPossibleInfected)
				}
			}
		}
	}

	return clientsContactInfected
}

func processContactPeople(infectedNotAnalysed map[int]models.InfectedSpreadPeriod,
	clientsContactInfected map[int][]*models.MonitorClients, virusSpread models.VirusSpread) {

	for id, clientContactInfected := range clientsContactInfected {

		clientsContactInfectedslice := make(models.MonitorClientsSlice, 0, len(clientContactInfected))
		for _, d := range clientContactInfected {
			clientsContactInfectedslice = append(clientsContactInfectedslice, *d)
		}

		sort.Sort(clientsContactInfectedslice)

		firstContact := clientsContactInfectedslice[0].Time
		lastContact := clientsContactInfectedslice[len(clientContactInfected)-1].Time

		spreadInitial := firstContact.AddDate(0, 0, -virusSpread.InfectionPeriod)
		spreadFinal := lastContact.AddDate(0, 0, virusSpread.InfectionPeriod)

		if _, ok := infectedNotAnalysed[id]; ok {
			previousFirstContact := infectedNotAnalysed[id].InfectionPeriod[0]
			previousLastContact := infectedNotAnalysed[id].InfectionPeriod[1]

			if previousFirstContact.Sub(firstContact).Minutes() < 0.0 {
				spreadInitial = previousFirstContact
			}

			if previousLastContact.Sub(spreadFinal).Minutes() > 0.0 {
				spreadFinal = previousLastContact
			}
		}

		virusSpreadPeriod := [2]time.Time{spreadInitial, spreadFinal}

		infectedPerson := models.InfectedSpreadPeriod{
			IDclient:        id,
			InfectionPeriod: virusSpreadPeriod,
		}

		infectedNotAnalysed[id] = infectedPerson
	}
}
