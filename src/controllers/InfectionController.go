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

var userExist = true

// FindInfectedController is a controller for find infected people
type FindInfectedController struct {
	controller.BaseController
	Routes []string
}

// Home login home page
func (t *FindInfectedController) Home() {

	if !userExist {
		t.Ctx.Data["UserExist"] = "User not found."
	}

	t.Ctx.Template = "homelogin"
	t.HTML(http.StatusOK)
}

// GetSignInUser login home page
func (t *FindInfectedController) GetSignInUser() {
	t.Ctx.Template = "registeruser"
	t.HTML(http.StatusOK)
}

// PostSignInUser login home page
func (t *FindInfectedController) PostSignInUser() {
	client := &models.Clients{}
	req := t.Ctx.Request()
	_ = req.ParseForm()
	if err := decoder.Decode(client, req.PostForm); err != nil {
		t.Ctx.Data["Message"] = err.Error()
		t.Ctx.Template = "error"
		t.HTML(http.StatusInternalServerError)
		return
	}

	t.Ctx.DB.Create(client)

	pathRedirect := fmt.Sprintf("/homeinfected/user/%d", client.ID)
	t.Ctx.Redirect(pathRedirect, http.StatusFound)
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
	strID := t.Ctx.Params["id"]
	id, err := strconv.Atoi(strID)
	if err != nil {
		t.Ctx.Data["Message"] = err.Error()
		t.Ctx.Template = "error"
		t.HTML(http.StatusInternalServerError)
		return
	}

	clients := []*models.Clients{}
	t.Ctx.DB.Where("id = ?", id).Find(&clients)

	if len(clients) == 0 {
		userExist = false
		t.Ctx.Redirect("/homeinfected", http.StatusFound)
		return
	}
	userExist = true

	t.Ctx.Data["ClientData"] = clients
	t.Ctx.Template = "userpage"
	t.HTML(http.StatusOK)
}

//UserPageTrackingData show the tracked positions of a given client
func (t *FindInfectedController) UserPageTrackingData() {
	strclientID := t.Ctx.Params["id"]
	clientID, err := strconv.Atoi(strclientID)
	if err != nil {
		t.Ctx.Data["Message"] = err.Error()
		t.Ctx.Template = "error"
		t.HTML(http.StatusInternalServerError)
		return
	}

	monitoredClient := []*models.MonitorClients{}
	t.Ctx.DB.Where("idclient = ?", clientID).Find(&monitoredClient)

	t.Ctx.Data["List"] = monitoredClient
	t.Ctx.Template = "userpageTrackingData"
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

	type Position struct {
		UserID    int     `json:"user"`
		Latitude  float32 `json:"lati"`
		Longitude float32 `json:"long"`
	}

	data := Position{}

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&data)
	if err != nil {
		log.Fatalln(err)
	}

	monitoredClient := &models.MonitorClients{}
	monitoredClient.IDclient = data.UserID
	monitoredClient.Latitude = data.Latitude
	monitoredClient.Longitude = data.Longitude
	monitoredClient.Time = time.Now()

	t.Ctx.DB.Create(monitoredClient)
}

// "get;/registeruser;GetSignInUser",
// "post;/postregisteruser;PostSignInUser",

//NewFindInfectedControllerController returns a new FindInfectedController
// (get or post); url ; method
func NewFindInfectedControllerController() controller.Controller {
	return &FindInfectedController{
		Routes: []string{
			"get;/homeinfected;Home",
			"get;/registeruser;GetSignInUser",
			"post;/postregisteruser;PostSignInUser",
			"post;/finduser;LogIn",
			"get;/homeinfected/user/{id};UserPage",
			"get;/homeinfected/user/trackposition/{id};UserPageTrackingData",
			"get;/homeinfected/user/declareinfection/{id};DeclareInfectionHome",
			"post;/homeinfected/user/declareinfection/postInfectionDate/{id};DeclareInfection",
			"post;/usercoordinates;UserCoordinates",
		},
	}
}
