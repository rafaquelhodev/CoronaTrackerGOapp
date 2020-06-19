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

//NewAdminController is the controller for the admin
// (get or post); url ; method
func NewAdminController() controller.Controller {
	return &AdminController{
		Routes: []string{
			"get;/admin;Home",
			"get;/admin/addinfected;DeclareInfection",
			"post;/admin/adminpostinfected;PostDeclareInfection",
		},
	}
}
