package controllers

import (
	"models"
	"net/http"
	"strconv"

	"github.com/gernest/utron/controller"
	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

// MediumController is a controller for Users list
type MediumController struct {
	controller.BaseController
	Routes []string
}

// Home renders a Users list
func (t *MediumController) Home() {
	users := []*models.Users{}
	t.Ctx.DB.Order("created_at desc").Find(&users)
	t.Ctx.Data["List"] = users
	t.Ctx.Template = "index"
	t.HTML(http.StatusOK)
}

//Create creates a User item
func (t *MediumController) Create() {
	users := &models.Users{}
	req := t.Ctx.Request()
	_ = req.ParseForm()
	if err := decoder.Decode(users, req.PostForm); err != nil {
		t.Ctx.Data["Message"] = err.Error()
		t.Ctx.Template = "error"
		t.HTML(http.StatusInternalServerError)
		return
	}

	t.Ctx.DB.Create(users)
	t.Ctx.Redirect("/", http.StatusFound)
}

//Delete deletes a User item
func (t *MediumController) Delete() {
	userID := t.Ctx.Params["id"]
	ID, err := strconv.Atoi(userID)
	if err != nil {
		t.Ctx.Data["Message"] = err.Error()
		t.Ctx.Template = "error"
		t.HTML(http.StatusInternalServerError)
		return
	}
	t.Ctx.DB.Delete(&models.Users{ID: ID})
	t.Ctx.Redirect("/", http.StatusFound)
}

//NewMediumController returns a new User list controller
// (get or post); url ; method
func NewMediumController() controller.Controller {
	return &MediumController{
		Routes: []string{
			"get;/;Home",
			"post;/create;Create",
			"get;/delete/{id};Delete",
		},
	}
}
