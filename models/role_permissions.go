package models

import (
	//"errors"
	"github.com/astaxie/beego"
)

type RolePermissions struct {
	Role_unique_id       int    `json:"role_unique_id"`
	Permission_unique_id string `json:"permission_unique_id"`
	Is_deleted           string `json:"permission_name"`
	Update_time          string `json:"update_time"`
	Create_time          string `json:"create_time"`
}

func (u *SysUser) RolePermissionsUpdate() string {
	return beego.AppConfig.String("rbac_sys_user_table")
}
func (u *SysUser) RolePermissionsCreat() string {
	return beego.AppConfig.String("rbac_sys_user_table")
}
