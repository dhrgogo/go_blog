package models

import (
	//"errors"
	"github.com/astaxie/beego"
)

type Role struct {
	Id             int    `json:"id"`
	Role_unique_id string `json:"role_unique_id"`
	Role_name      string `json:"role_name"`
	RoleDesc       string `json:"roleDesc"`
	Is_deleted     string `json:"is_deleted"`
	Update_time    string `json:"update_time"`
	Create_time    string `json:"create_time"`
}

func (u *SysUser) RoleById() string {
	return beego.AppConfig.String("rbac_sys_user_table")
}
func (u *SysUser) RolerList() string {
	return beego.AppConfig.String("rbac_sys_user_table")
}
func (u *SysUser) RoleUpdateById() string {
	return beego.AppConfig.String("rbac_sys_user_table")
}
func (u *SysUser) RoleUpdateByUser_unique_id() string {
	return beego.AppConfig.String("rbac_sys_user_table")
}
func (u *SysUser) RoleCreat() string {
	return beego.AppConfig.String("rbac_sys_user_table")
}
