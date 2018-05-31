package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	// "fmt"
)
type User struct {
	Id        int `orm:"pk;auto"`
	Username  string `orm:"unique"`
	Password  string
	Token     string `orm:"unique"`
	Avatar    string
	Email     string `orm:"null"`
	Url       string `orm:"null"`
	Signature string `orm:"null;size(1000)"`
	InTime    time.Time `orm:"auto_now_add;type(datetime)"`
}
func SaveUser(user *User) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(user)
	return id
}
func UpdateUser(user *User) {
	o := orm.NewOrm()
	o.Update(user)
}
func DeleteUser(user *User) {
    o := orm.NewOrm()
    o.Delete(user)
}
func FindUserById(id int) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Id", id).One(&user)
	return err != orm.ErrNoRows, user
}
func FindUserByToken(token string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Token", token).One(&user)
	return err != orm.ErrNoRows, user
}
func FindUserByUserName(username string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Username", username).One(&user)
	return err != orm.ErrNoRows, user
}