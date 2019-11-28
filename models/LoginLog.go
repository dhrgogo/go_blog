package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type LoginLog struct {
	Id         int64     `orm:"auto"`
	Username   string    `orm:"size(100)" form:"Username" valid:"Required;MaxSize(20);MinSize(1)"`
	Createtime time.Time `orm:"type(datetime);auto_now_add"`
}

func (u *LoginLog) TableName() string {
	return beego.AppConfig.String("rbac_login_log_table")
}
func init() {
	orm.RegisterModel(new(LoginLog))
}
