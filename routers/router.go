package routers

import (
	"eat/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/eat", &controllers.EatController{})
	beego.Router("/user/eat", &controllers.UserController{}, "*:Eat")
	beego.Router("/login/code", &controllers.UserController{}, "*:Code")
}
