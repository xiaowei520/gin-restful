package model

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"log"

	. "../conf"

)

const (
	timeFormate = "2006-01-02 15:04:05"
)

//声明 数据库变量
var sqlDB *sql.DB


func init(){
	Open()
}
/*
连接数据库
 */
func Open() (*sql.DB) {
	if sqlDB == nil {
		log.Println("start new db connect")

		sqlConnection := Conf.DB.UserName + ":" + Conf.DB.Pwd + "@tcp(" + Conf.DB.Host + ":" + Conf.DB.Port + ")/" + Conf.DB.Name + "?charset=utf8mb4&parseTime=True"

		//sql.Open方法会创建一个数据库连接池db,这个db不是数据库连接，它是一个连接池，只有当真正数据库通信的时候才创建连接
		db, err := sql.Open("mysql", sqlConnection)
	//	defer db.Close() //遇到错误-连接数据库关闭
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		sqlDB = db
		//压力测试 如果不设置 存在大量TIME_WAIT 连接
		sqlDB.SetMaxIdleConns(10)  //设置数据库的空闲连接
		sqlDB.SetMaxOpenConns(100) //设置最大打开连接

		//列如sqlDB.ping 并不创建真正连接
		if err := sqlDB.Ping(); err != nil {
			log.Fatalln(err)
		}

	}
	log.Println(sqlDB)
	return sqlDB
}