package controllers

import (
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
	// "newApp/models"
	// "fmt"
)

type MainController struct {
	beego.Controller
}
type Mler struct {
	beego.Controller
}
type JSONStruct struct {
    Code int
    Msg  string
}
// func (c *MainController) Get() {
// 	c.Data["Website"] = "beego.me"
// 	c.Data["Email"] = "astaxie@gmail.com"
// 	c.TplName = "index.html"
// }
func (c *MainController) Get() {
	c.Data["Website"] = "董化仁"
	c.Data["Email"] = "616011986@qq.com"
	c.TplName = "index.html"
}

func (c *Mler) Get() {
	c.Data["Website"] = "董化仁"
	c.Data["Email"] = "616011986@qq.com"
	c.TplName = "index.html"
}
