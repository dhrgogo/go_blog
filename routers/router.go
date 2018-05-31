package routers

import (
	"go_blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/test", &controllers.Mler{})
	// beego.Router("/user", &controllers.user{},"POST:Login")
	beego.Router("/user", &controllers.UserControllers{}, "GET:Search")
	beego.Router("/user", &controllers.UserControllers{}, "POST:Creat")
    beego.Router("/user", &controllers.UserControllers{}, "PUT:Update")
    beego.Router("/user", &controllers.UserControllers{}, "DELETE:Delete")
}
