package models

import (
	//"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type SysUser struct {
	Id             int       `orm:"pk;auto;column(id);description(唯一id)"`
	User_unique_id string    `orm:"unique;size(32);column(user_unique_id);description(唯一id)"`
	Name           string    `orm:"size(32);column(name);description(名字)"`
	Phone          string    `orm:"size(32);column(phone);description(手机号)"`
	Password       string    `orm:"size(32);column(password);description(密码)"`
	Post           string    `orm:"size(32);column(post);description(职位)"`
	English_name   string    `orm:"size(32);column(english_name);description(英文名称)"`
	Department     string    `orm:"size(100);column(department);description(部门)"`
	Reorder        string    `orm:"size(11);column(reorder);description(自定义排序)"`
	Sex            int       `orm:"size(4);column(sex);description(0女、1男、2未知)"`
	Email          string    `orm:"size(80);column(email);description(邮箱)"`
	Header_img     string    `orm:"size(255);column(header_img);description(头像)"`
	Is_deleted     int       `orm:"size(4);column(is_deleted);default(1);description(0删除、1未删除)"`
	Last_name      string    `orm:"size(32);column(last_name);description(姓氏)" `
	Nickname       string    `orm:"size(32);column(nickname);description(昵称)"`
	Openid         string    `orm:"size(32);column(openid);description(第三方登录唯一id)" `
	Status         int       `orm:"size(4);column(status);default(0)description(0:未激活,1:已激活,2:待审核,3:已拒绝)"`
	Create_time    time.Time `orm:"type(datetime);auto_now_add" `
	Update_time    time.Time `orm:"type(datetime);auto_now"`
	//Role                  []*Role   `orm:"rel(m2m)"`
}

func (u *SysUser) TableName() string {
	return beego.AppConfig.String("rbac_sys_user_table")
}
func init() {
	orm.RegisterModel(new(SysUser))
}
