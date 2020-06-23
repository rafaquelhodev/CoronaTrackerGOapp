package controllers

import (
	"fmt"
	"models"
	"net/http"
	"strconv"
	"time"

	"github.com/gernest/utron/controller"
)

// AdminController is a controller for the admin of the APP
type AdminController struct {
	controller.BaseController
	Routes []string
}

// Home is the admin home page
func (t *AdminController) Home() {
	t.Ctx.Template = "homeadmin"
	t.HTML(http.StatusOK)
}

// DeclareInfection logs in at the admin home page
func (t *AdminController) DeclareInfection() {
	t.Ctx.Template = "admindeclareinfection"
	t.HTML(http.StatusOK)
}

//PostDeclareInfection POST -> declare infection date of a specific client
func (t *AdminController) PostDeclareInfection() {
	req := t.Ctx.Request()

	date := req.FormValue("TestingDate")
	dateParsed, err := time.Parse("2006-01-02", date)
	if err != nil {
		fmt.Println("Error parsing time, Error = ", err.Error())
	}

	infected := &models.Infecteds{}
	infected.TestingDate = dateParsed
	infected.IDclient, err = strconv.Atoi(req.FormValue("IDclient"))
	if err != nil {
		fmt.Println("Error getting client ID, Error = ", err.Error())
	}
	t.Ctx.DB.Create(infected)

	pathRedirect := "/admin"
	t.Ctx.Redirect(pathRedirect, http.StatusFound)
}

// FindInfected : POST method that finds infected people
func (t *AdminController) FindInfected() {

	virusSpread := models.VirusSpread{}
	virusSpread.InfectionPeriod = 20
	virusSpread.ContactDistanceMeters = 5.0
	virusSpread.ContactTimeMinutes = 5.0

	// map with infected people that were not analysed yet
	infectedNotAnalysed := make(map[int]models.InfectedSpreadPeriod, 0)

	// array of analysed people that were infected
	infectedAnalysed := make(map[int]models.InfectedSpreadPeriod, 0)

	// Reading infected table and populating infectedPeople map
	t.HandleOnInfecteds(infectedNotAnalysed, virusSpread)

	var mainloop bool
	mainloop = true
	for mainloop {
		idClient, initialSpreadDate, finalSpreadDate := FindOldestInfected(infectedNotAnalysed)

		// Deleting for the future iterations
		delete(infectedNotAnalysed, idClient)

		// Populating map of infected clients that were already analysed
		infectedAnalysed[idClient] = models.InfectedSpreadPeriod{
			IDclient:        idClient,
			InfectionPeriod: [2]time.Time{initialSpreadDate, finalSpreadDate},
		}

		infectedData, mapAreasInfected := t.RetrieveDataInfectedPerson(idClient, initialSpreadDate, finalSpreadDate)

		possibleInfecteds := t.CommonAreaDate(initialSpreadDate, finalSpreadDate, mapAreasInfected, infectedAnalysed)

		contactWithInfected := contactWithInfected(infectedData, possibleInfecteds, virusSpread)

		processContactPeople(infectedNotAnalysed, contactWithInfected, virusSpread)

		if len(infectedNotAnalysed) == 0 {
			mainloop = false
		}
	}

	t.Ctx.Data["Infecteds"] = infectedAnalysed
	t.Ctx.Template = "searchresults"
	t.HTML(http.StatusOK)
}

//NewAdminController is the controller for the admin
// (get or post); url ; method
func NewAdminController() controller.Controller {
	return &AdminController{
		Routes: []string{
			"get;/admin;Home",
			"get;/admin/addinfected;DeclareInfection",
			"post;/admin/adminpostinfected;PostDeclareInfection",
			"get;/admin/adminrunsearch;FindInfected",
		},
	}
}
