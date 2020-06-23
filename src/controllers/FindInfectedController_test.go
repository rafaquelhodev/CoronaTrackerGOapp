package controllers

import (
	"models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gernest/utron/base"
	"github.com/gernest/utron/controller"
	utronModels "github.com/gernest/utron/models"
	"github.com/jinzhu/gorm"
)

func TestDistance(t *testing.T) {
	var lati1, lati2, long1, long2 float32
	lati1 = 5.0
	long1 = 0.0
	lati2 = 5.0
	long2 = 0.0

	dist := distance(lati1, long1, lati2, long2)

	if dist != 0.0 {
		t.Errorf("Expected zero distance; got %f", dist)
	}
}

func TestHandleOnInfecteds(t *testing.T) {

	// Arrange

	// open gorm db using mocked database
	gdb, err := gorm.Open("sqlite3", "Teste.db")
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	// assigning the test database to the controller:
	modelUltron := utronModels.NewModel()
	modelUltron.DB = gdb

	modelUltron.Register(&models.Infecteds{})
	modelUltron.AutoMigrateAll()

	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	Ctx := base.NewContext(w, req)
	Ctx.DB = modelUltron

	baseController := controller.BaseController{
		Ctx: Ctx,
	}

	admController := AdminController{}
	admController.BaseController = baseController
	//--------------------------------------------

	infected := &models.Infecteds{}
	admController.Ctx.DB.Delete(infected)

	infected.ID = 1
	infected.TestingDate = time.Date(2020, 6, 21, 0, 0, 0, 0, time.UTC)
	infected.IDclient = 1
	admController.Ctx.DB.Create(infected)

	infected.ID = 3
	infected.TestingDate = time.Date(2020, 6, 22, 0, 0, 0, 0, time.UTC)
	infected.IDclient = 3
	admController.Ctx.DB.Create(infected)

	infectedNotAnalysed := make(map[int]models.InfectedSpreadPeriod, 0)

	virusSpread := models.VirusSpread{}
	virusSpread.InfectionPeriod = 20
	virusSpread.ContactDistanceMeters = 5.0
	virusSpread.ContactTimeMinutes = 5.0

	// Act
	admController.HandleOnInfecteds(infectedNotAnalysed, virusSpread)

	// Assert
	if len(infectedNotAnalysed) != 2 {
		t.Errorf("Expected list of size = %d; got %d", 2, len(infectedNotAnalysed))
	}
}
