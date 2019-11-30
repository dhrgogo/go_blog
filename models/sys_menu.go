package models

import (
	//"errors"
	"github.com/astaxie/beego"
)

type SysMenu struct {
	Id             int    `json:"id"`
	Role_unique_id string `json:"role_unique_id"`
	Menu           string `json:"menu"`
	Url            string `json:"url"`
	Sort           string `json:"sort"`
	Style          string `json:"style"`
	Parent_id      string `json:"parent_id"`
	Is_deleted     string `json:"is_deleted"`
}

func (u *SysUser) SysMenuById() string {
	return beego.AppConfig.String("rbac_sys_user_table")
}
func (u *SysUser) SysMenuList() string {
	return beego.AppConfig.String("rbac_sys_user_table")
}
func (u *SysUser) SysMenuUpdateById() string {
	return beego.AppConfig.String("rbac_sys_user_table")
}
func (u *SysUser) SysMenuUpdateByUser_unique_id() string {
	return beego.AppConfig.String("rbac_sys_user_table")
}
func (u *SysUser) SysMenuCreat() string {
	return beego.AppConfig.String("rbac_sys_user_table")
}
