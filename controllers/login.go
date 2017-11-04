package controllers

import (
	"dataStore/models"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"
	//var err error
	var err error
	this.Data["users"], err = models.GetUser()

	if err != nil {
		beego.Error(err)
	}
}

func (this *LoginController) Post() {
	//account := this.Input().Get("account")
	//password := this.Input().Get("password")
	//autoLogin := this.Input().Get("autoLogin") == "on"
	//this.Ctx.WriteString(account + password)
	this.Redirect("/", 301)
	return
}
