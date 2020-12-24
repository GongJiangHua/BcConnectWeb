package main

import (
	"BitConnectWeb/db_mysql"
	_ "BitConnectWeb/routers"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"

)

func main() {
	//连接数据库
	db_mysql.ConnectDb()
	//设置静态资源文件映射
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/img","./static/img")
	beego.SetStaticPath("/css","./static/css")
	beego.Run()

}

