package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"models"
	"net/http"
	"strconv"
	"time"

	"github.com/gernest/utron/controller"
)

// FindInfectedController is a controller for find infected people
type FindInfectedController struct {
	controller.BaseController
	Routes []string
}

// Home login home page
func (t *FindInfectedController) Home() {
	t.Ctx.Template = "homelogin"
	t.HTML(http.StatusOK)
}

//LogIn a specific client
func (t *FindInfectedController) LogIn() {
	req := t.Ctx.Request()
	id := req.FormValue("Id")

	pathRedirect := fmt.Sprintf("/homeinfected/user/%s", id)
	t.Ctx.Redirect(pathRedirect, http.StatusFound)
}

//UserPage home page of a specific client
func (t *FindInfectedController) UserPage() {
	strclientID := t.Ctx.Params["id"]
	clientID, err := strconv.Atoi(strclientID)
	if err != nil {
		t.Ctx.Data["Message"] = err.Error()
		t.Ctx.Template = "error"
		t.HTML(http.StatusInternalServerError)
		return
	}

	clients := []*models.Clients{}
	t.Ctx.DB.Where("idclient = ?", clientID).Find(&clients)
	t.Ctx.Data["List"] = clients
	t.Ctx.Template = "userpage"
	t.HTML(http.StatusOK)
}

//DeclareInfectionHome GET -> declare infection date of a specific client
func (t *FindInfectedController) DeclareInfectionHome() {
	t.Ctx.Template = "declareinfection"
	t.HTML(http.StatusOK)
}

//DeclareInfection POST -> declare infection date of a specific client
func (t *FindInfectedController) DeclareInfection() {
	strclientID := t.Ctx.Params["id"]
	clientID, err := strconv.Atoi(strclientID)
	if err != nil {
		t.Ctx.Data["Message"] = err.Error()
		t.Ctx.Template = "error"
		t.HTML(http.StatusInternalServerError)
		return
	}

	infected := &models.Infecteds{}
	req := t.Ctx.Request()
	_ = req.ParseForm()

	date := req.FormValue("testing_date")
	dateParsed, err := time.Parse("2006-01-02", date)

	if err != nil {
		t.Ctx.Data["Message"] = err.Error()
		t.Ctx.Template = "error"
		t.HTML(http.StatusInternalServerError)
		return
	}
	infected.TestingDate = dateParsed
	infected.IDclient = clientID
	t.Ctx.DB.Create(infected)

	pathRedirect := fmt.Sprintf("/homeinfected/user/%d", clientID)
	t.Ctx.Redirect(pathRedirect, http.StatusFound)
}

//UserCoordinates of a specific client
func (t *FindInfectedController) UserCoordinates() {
	req := t.Ctx.Request()
	decoder := json.NewDecoder(req.Body)

	// corpo, err := ioutil.ReadAll(req.Body)
	// if err != nil {
	// 	fmt.Println("[main] erro ao ler servidor. Erro: ",
	// 		err.Error())
	// 	return
	// }
	// fmt.Println(" ")

	type Data struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	}

	type Position struct {
		User int     `json:"user"`
		Lati float32 `json:"lati"`
		Long float32 `json:"long"`
	}

	data := Position{}
	err := decoder.Decode(&data)
	if err != nil {
		log.Fatalln(err)
	}

	// err = json.Unmarshal(corpo, &data)
	fmt.Printf("%v", data)

	// lati := req.FormValue("latitude")
	// long := req.FormValue("longitude")
	// fmt.Println("Latitude = ", lati, "; longitude = ", long)

	// pathRedirect := fmt.Sprintf("/homeinfected/user/%s", strclientID)
	// t.Ctx.Redirect(pathRedirect, http.StatusFound)
}

//NewFindInfectedControllerController returns a new FindInfectedController
// (get or post); url ; method
func NewFindInfectedControllerController() controller.Controller {
	return &FindInfectedController{
		Routes: []string{
			"get;/homeinfected;Home",
			"post;/finduser;LogIn",
			"get;/homeinfected/user/{id};UserPage",
			"get;/homeinfected/user/declareinfection/{id};DeclareInfectionHome",
			"post;/homeinfected/user/declareinfection/postInfectionDate/{id};DeclareInfection",
			"post;/usercoordinates;UserCoordinates",
		},
	}
}
