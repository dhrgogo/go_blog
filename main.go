package main

import (
	_ "admin/models"
	_ "admin/routers"
	"fmt"
	"github.com/astaxie/beego"
)

const VERSION = "0.1.2"

//func initialize() {
//	mime.AddExtensionType(".css", "text/css")
//	//判断初始化参数
//	initArgs()
//	models.Connect()
//}
func initArgs() {
	//models.Syncdb()
}

func main() {
	//初始化
	//initialize()
	fmt.Println("Starting....")

	fmt.Println("Start ok")
	// 开启 ORM 调试模式
	//orm.Debug = true
	// 自动建表
	//orm.RunSyncdb("default", true, true)
	//beego.AddFuncMap("stringsToJson", models.StringsToJson)

	// 运行时
	//models.Syncdb()

	beego.Run()
}
