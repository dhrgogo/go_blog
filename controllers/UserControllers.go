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
	form := models.User{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	models.SaveUser(&form)
	c.Data["json"] = map[string]interface{}{"ErrCode": 0, "Result": &form}
	c.ServeJSON()
}
//查询用户 
func (c *UserControllers) Search() {
	token := c.GetString("token")
	if token == "" {
		c.Ctx.WriteString("token is empty");
		return
	}
	ok, user := models.FindUserByToken(token)
	// c.Ctx.WriteString("Search")
	mystruct := user
    c.Data["json"] = map[string]interface{}{"ErrCode": 0, "Result": mystruct}
	c.ServeJSON()
	fmt.Println("Search",ok,user)
}
//更新用户
func (c *UserControllers) Update() {
	form := models.User{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	models.UpdateUser(&form)
	c.Data["json"] = map[string]interface{}{"ErrCode": 0, "Result": &form}
	c.ServeJSON()
}
//删除用户
func (c *UserControllers) Delete() {
	c.Data["json"] = map[string]interface{}{"success": 0, "message": "111"}
	c.ServeJSON()
}