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
func (c *MainController) Get() {
	c.Data["Website"] = "dhr"
	c.Data["Email"] = "123XXXXXX@qq.com"
	c.TplName = "index.html"
}

func (c *Mler) Get() {
	c.Data["Website"] = "dhr"
	c.Data["Email"] = "123XXXXXX@qq.com"
	c.TplName = "index.html"
}
