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

// Home login home page
func (t *AdminController) Home() {
	t.Ctx.Template = "homeadmin"
	t.HTML(http.StatusOK)
}

//NewFindInfectedControllerController returns a new FindInfectedController
// (get or post); url ; method
func NewAdminController() controller.Controller {
	return &AdminController{
		Routes: []string{
			"get;/admin;Home",
		},
	}
}
