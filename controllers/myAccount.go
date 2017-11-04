package controllers

import (
	"github.com/astaxie/beego"
)

type MyAccountController struct {
	beego.Controller
}

func (this *MyAccountController) Get() {
	this.TplName = "myAccount.html"
}
