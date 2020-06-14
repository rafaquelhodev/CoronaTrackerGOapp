package controllers

import (
	"fmt"
	"models"
	"net/http"
	"strconv"

	"github.com/gernest/utron/controller"
)

// FindInfectedController is a controller for find infected people
type FindInfectedController struct {
	controller.BaseController
	Routes []string
}

// Home login home page
func (t *FindInfectedController) Home() {
	t.Ctx.Template = "homeinfected"
	t.HTML(http.StatusOK)
}

//LogIn a specific client
func (t *FindInfectedController) LogIn() {
	req := t.Ctx.Request()
	id := req.FormValue("Id")

	pathRedirect := fmt.Sprintf("/user/%s", id)
	t.Ctx.Redirect(pathRedirect, http.StatusFound)
}

//HomeUser home page of a specific client
func (t *FindInfectedController) HomeUser() {
	// req := t.Ctx.Request()

	clientID := t.Ctx.Params["id"]
	ID, err := strconv.Atoi(clientID)
	if err != nil {
		t.Ctx.Data["Message"] = err.Error()
		t.Ctx.Template = "error"
		t.HTML(http.StatusInternalServerError)
		return
	}

	fmt.Println("user id = ", ID)

	clients := []*models.Clients{}
	t.Ctx.DB.Find(&clients).Find(&models.Clients{IDclient: ID})
	t.Ctx.Data["List"] = clients
	t.Ctx.Template = "homeuser"
	t.HTML(http.StatusOK)
}

//NewFindInfectedControllerController returns a new FindInfectedController
// (get or post); url ; method
func NewFindInfectedControllerController() controller.Controller {
	return &FindInfectedController{
		Routes: []string{
			"get;/homeinfected;Home",
			"post;/finduser;LogIn",
			"get;/user/{id};HomeUser",
		},
	}
}
