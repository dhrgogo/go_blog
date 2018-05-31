package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"go_blog/models"
	"fmt"
)
type UserControllers struct {
	beego.Controller
}

//创建用户
func (c *UserControllers ) Creat() {
	// user := models.User{
	// 	Username:"dhr",
	// 	Password:"616011986",
	// 	Token:"110110",
	// 	}
	// models.SaveUser(&user)
	form := models.User{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	c.Ctx.WriteString("Creat")
}
//查询用户 
func (c *UserControllers) Search() {
	ok, user := models.FindUserByToken("110110")
	// c.Ctx.WriteString("Search")
	mystruct := user
    c.Data["json"] = map[string]interface{}{"ErrCode": 0, "Result": mystruct}
	c.ServeJSON()
	fmt.Println("Search",ok,user)
}
//更新用户
func (c *UserControllers) Update() {
	c.Ctx.WriteString("Update")
	fmt.Println("Update")
}
//删除用户
func (c *UserControllers) Delete() {
	c.Data["json"] = map[string]interface{}{"success": 0, "message": "111"}
	c.ServeJSON()
}