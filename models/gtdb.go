package models

//
//import (
//	"errors"
//	logger "log"
//	"time"
//	"xorm.io/core"
//
//	"github.com/astaxie/beego"
//	"github.com/go-xorm/xorm"
//	uuid "github.com/satori/go.uuid"
//	// "os"
//	//"encoding/json"
//)
//
//var X, x0, x1 *xorm.Engine  //定义引擎全局变量(分别表示当前使用的引擎，模拟数据库引擎，生产数据库引擎)
//var Cur_db string           // 当前使用的数据库，为 "dev" / "prod" ,分别代表模拟数据库和生产数据库
//var AllTables []interface{} // 存储所有表的struct结构
//
//// 配置表`
//type Configs struct {
//	Conf map[string]string `json:"conf"`                   // 具体的配置信息
//	Tag  string            `xorm:"varchar(50)" json:"tag"` // 标签，用于区别不同模式的配置
//}
//
//
//
//
//// 角色表
//type Role struct {
//	Id        int      `xorm:"autoincr pk index" json:"id,omitempty"`     // 角色ID，超级角色super的ID一定是1
//	Father_id int      `xorm:" notnull" json:"father_id,omitempty"`       // 角色的上级角色ID，若为超级用户角色，则此字段为 -1
//	Uid       string   `xorm:"varchar(36) notnull" json:"uid,omitempty"`  // 角色的上级用户ID，若为超级角色，则此字段为 '1'
//	Name      string   `xorm:"varchar(35) notnull" json:"name,omitempty"` // 角色名称
//	Des       string   `xorm:"varchar(350)" json:"des,omitempty"`         // 角色描述
//	Funcs     []string `json:"funcs,omitempty"`                           // 可用的功能编号列表
//	Pwd_login bool     `xorm:"default(true)" json:"pwd_login,omitempty"`  // 是否允许密码登录（标注员角色的值为false）
//
//	// {bug_alarm(yes/no) => 是否开启故障报警邮件推送}
//	Extra map[string]string `json:"extra,omitempty"`               // 额外信息，包含信息{ bug_alarm=>yes/no , ...}
//	Time  int               `xorm:"created" json:"time,omitempty"` // 角色创建时间
//}
//
//
//// 启动程序后就执行
//func init() {
//	// 创建 ORM 引擎与数据库
//	var err error
//	x0, err = xorm.NewEngine("postgres", beego.AppConfig.String("dburi"))
//	if err != nil {
//		logger.Panicln("模拟数据库 引擎创建出错: ", err)
//		return
//	}
//	x1, err = xorm.NewEngine("postgres", beego.AppConfig.String("dburi_prod"))
//	if err != nil {
//		logger.Panicln("生产数据库 引擎创建出错: ", err)
//		return
//	}
//
//	AllTables = append(AllTables,
//		// 用户相关
//		new(User),           // 用户表
//		new(Role),          // 角色表
//
//		// 其他表
//		new(Configs), // 全局性配置表
//	)
//
//	err = initDb(x0)
//	if err != nil {
//		logger.Panicln("模拟数据库 初始化出错: ", err)
//		return
//	}
//	err = initDb(x1)
//	if err != nil {
//		logger.Panicln("生产数据库 初始化出错: ", err)
//		return
//	}
//
//	Use_dev_db()
//}
//
//func initDb(X *xorm.Engine) error {
//	err := X.Ping()
//	if err != nil {
//		return errors.New("Ping 数据库失败: " + err.Error())
//	}
//	X.SetConnMaxLifetime(60 * time.Second)
//	X.SetMaxIdleConns(40)
//	X.SetMaxOpenConns(100)
//
//	go func() {
//		for { // 每隔60秒 ping 一次数据库
//			time.Sleep(180 * time.Second)
//			X.Ping()
//		}
//	}()
//
//	// // 打印具体日志
//	X.ShowSQL(true)
//	X.Logger().SetLevel(core.LOG_DEBUG)
//
//	// 同步结构体与数据表,和python的 migrate 一样同步数据库
//	err = X.Sync2(AllTables...)
//	if err != nil {
//		return errors.New("同步数据库失败: " + err.Error())
//	}
//
//	cacher := xorm.NewLRUCacher2(xorm.NewMemoryStore(), time.Duration(300)*time.Second, 5000)
//	X.MapCacher(new(Dev_type), cacher)
//	X.MapCacher(new(Bug_type), cacher)
//	X.MapCacher(new(Configs), cacher)
//	X.MapCacher(new(Role), cacher)
//	X.MapCacher(new(Device), cacher)
//	return nil
//}
//
//// 使用 模拟数据库
//func Use_dev_db() {
//	X = x0
//	Cur_db = "dev"
//}
//
//// 使用 生产数据库
//func Use_prod_db() {
//	X = x1
//	Cur_db = "prod"
//}
//
////// 同步扩展功能的表模型
////func Sync2Ext(tables []interface{}) error {
////	err := x0.Sync2(tables...)
////	if err != nil {
////		return errors.New("同步演示数据库失败: " + err.Error())
////	}
////	err = x1.Sync2(tables...)
////	if err != nil {
////		return errors.New("同步生产数据库失败: " + err.Error())
////	}
////	return nil
////}
