package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	// "fmt"
)
type Artical struct {
	Id        int `orm:"pk;auto"`
	author  string `orm:"unique"`
	countent  string
	title     string `orm:"unique"`
	InTime    time.Time `orm:"auto_now_add;type(datetime)"`
}
func SaveArtical(artical *Artical) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(artical)
	return id
}
func UpdateArtical(artical *Artical) {
	o := orm.NewOrm()
	o.Update(artical)
}
func DeleteArtical(artical *Artical) {
    o := orm.NewOrm()
    o.Delete(artical)
}