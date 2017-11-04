package controllers

import (
	"github.com/astaxie/beego"
)

type AddApplyController struct {
	beego.Controller
}

func (this *AddApplyController) Get() {
	this.TplName = "addApplay.html"
}
