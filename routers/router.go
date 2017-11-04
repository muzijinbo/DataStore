package routers

import (
	"dataStore/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/myAccount", &controllers.MyAccountController{})
	beego.Router("/addApply", &controllers.AddApplyController{})
}
