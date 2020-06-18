package controllers

import (
	"net/http"

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

//NewAdminController is the controller for the admin
// (get or post); url ; method
func NewAdminController() controller.Controller {
	return &AdminController{
		Routes: []string{
			"get;/admin;Home",
			"get;/admin/addinfected;DeclareInfection",
		},
	}
}
