package models

import (
	//"errors"
	"github.com/astaxie/beego"
	"time"
)

type SysUser struct {
	Id             int       `json:"id"`
	User_unique_id string    `json:"user_unique_id"`
	Name           string    `json:"name"`
	Phone          string    `json:"phone"`
	Password       string    `json:"password"`
	Post           string    `json:"post"`
	English_name   string    `json:"english_name"`
	Department     string    `json:"department"`
	Reorder        string    `json:"reorder"`
	Sex            int       `json:"sex"`
	Email          string    `json:"email"`
	Header_img     string    `json:"header_img"`
	Is_deleted     int       `json:"is_deleted"`
	Last_name      string    `json:"last_name"`
	Nickname       string    `json:"nickname"`
	Openid         string    `json:"openid"`
	Status         int       `json:"Status"`
	Create_time    time.Time `json:"create_time"`
	Update_time    time.Time `json:"update_time"`
	//Role                  []*Role   `orm:"rel(m2m)"`
}

func (u *SysUser) SysUserById() string {
	return beego.AppConfig.String("rbac_sys_user_table")
}
func (u *SysUser) SysUserList() string {
	return beego.AppConfig.String("rbac_sys_user_table")
}
func (u *SysUser) SysUserUpdateById() string {
	return beego.AppConfig.String("rbac_sys_user_table")
}
func (u *SysUser) SysUserUpdateByUser_unique_id() string {
	return beego.AppConfig.String("rbac_sys_user_table")
}
func (u *SysUser) SysUserCreat() string {
	return beego.AppConfig.String("rbac_sys_user_table")
}
