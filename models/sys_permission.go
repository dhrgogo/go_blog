package models

import (
	//"errors"
	"github.com/astaxie/beego"
)

type SysPermission struct {
	Id                   int    `json:"role_unique_id"`
	Permission_unique_id string `json:"permission_unique_id"`
	Permission_name      string `json:"permission_name"`
	PermissionDesc       string `json:"permissionDesc"`
	Is_deleted           string `json:"is_deleted"`
	Update_time          string `json:"update_time"`
	Create_time          string `json:"create_time"`
	//Role                  []*Role   `orm:"rel(m2m)"`
}

func (u *SysUser) SysPermissionById() string {
	return beego.AppConfig.String("rbac_sys_user_table")
}
func (u *SysUser) SysPermissionList() string {
	return beego.AppConfig.String("rbac_sys_user_table")
}
func (u *SysUser) SysPermissionUpdateById() string {
	return beego.AppConfig.String("rbac_sys_user_table")
}
func (u *SysUser) SysPermissionUpdateByUser_unique_id() string {
	return beego.AppConfig.String("rbac_sys_user_table")
}
func (u *SysUser) SysPermissionCreat() string {
	return beego.AppConfig.String("rbac_sys_user_table")
}
