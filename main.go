package main

import (
    // "fmt"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "go_blog/models"
    _ "github.com/lib/pq"
    _ "go_blog/routers"
)

func init(){
    // PostgreSQL 配置
    orm.RegisterDriver("postgres", orm.DRPostgres) // 注册驱动
    orm.RegisterDataBase("default", "postgres", "user=postgres password=1234 dbname=test host=104.225.145.3 port=5432 sslmode=disable")
    orm.RegisterModel(
		new(models.User),
    )
    orm.RunSyncdb("default", false, true)
}

func main() {
    // orm.Debug = true
    // ok, err := regexp.MatchString("/topic/edit/[0-9]+", "/topic/edit/123")
    // beego.Debug(ok, err)
	beego.Run()
}

