package main

import (
	_ "beetwo2/routers"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	config:=beego.AppConfig

/*	name:=config.String("appname")
	fmt.Println(name)
	beego.Run()*/

	appName :=config.String("appname")
	fmt.Println("项目应用名称:",appName)
	port,err:=config.Int("httpport")
	if err != nil {
		panic("项目配置信息解释错误，请检验后重试")
	}
	fmt.Println("应用监听端口：",port)

	driver:=config.String("db_driver")
	dbUser:=config.String("dbUser")
	dbPassword:=config.String("dbPassword")
	dbIp:=config.String("dbIp")
	dbName:=config.String("dbName")

	db,err:=sql.Open(driver,dbUser+":"+dbPassword+"@tcp("+dbIp+")/login"+dbName+"?charset=utf8")
	if err != nil {
		panic("数据打开失败，请重试")
		fmt.Println(db)
	}
	beego.Run()
}

