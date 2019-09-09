package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	//数据库连接参数
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")

	//数据库连接字符串 dbConn := "root:123123@tcp(127.0.0.1:3306)/lemonblog?charset=utf8"
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	//注册数据库连接
	orm.RegisterDataBase("default", "mysql",  dbConn )

	// 注册模型
	//orm.RegisterModelWithPrefix("prefix_", new(User))
	orm.RegisterModel(new(User), new(Article), new(Comment), new(Category))

	// 创建数据库
	orm.RunSyncdb("default", false, true)

}