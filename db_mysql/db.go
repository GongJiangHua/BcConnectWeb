package db_mysql

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
)
var DB *sql.DB
func ConnectDb()  {
	//准备连接数据库要的数据
	config := beego.AppConfig
	dbDriver := config.String("db_driverName")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")
	fmt.Println(dbDriver,dbName,dbUser,dbPassword,dbIp)
	//连接数据库
	connUrl := dbUser +":" + dbPassword + "@tcp("+dbIp+")/"+dbName+"?charset=utf8"
	db,err := sql.Open(dbDriver,connUrl)
	if err != nil {
		//err不为nil，表示连接数据库出现错误，程序恐慌
		fmt.Println(err.Error())
		//panic("数据库连接错误，请检查配置")
	}
	fmt.Println(db)
	DB = db
}

