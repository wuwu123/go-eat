package routers

import (
	"eat/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/user",&controllers.UserController{})
	beego.Router("/admin",&controllers.AdminController{})
}
