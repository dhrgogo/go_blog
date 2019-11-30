package models

import (
	//"errors"
	"github.com/astaxie/beego"
)

type UserRole struct {
	Id             int    `json:"id"`
	User_unique_id string `json:"user_unique_id"`
	Role_unique_id string `json:"role_unique_id"`
	Is_deleted     string `json:"is_deleted"`
	Update_time    string `json:"update_time"`
	Create_time    string `json:"create_time"`
}

func (u *SysUser) UserRoleUpdate() string {
	return beego.AppConfig.String("rbac_sys_user_table")
}
func (u *SysUser) UserRoleCreat() string {
	return beego.AppConfig.String("rbac_sys_user_table")
}
