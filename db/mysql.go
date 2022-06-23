package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
	"ziweiSgServer/config"
)

var Engine *xorm.Engine

func TestDB() {
	mysqlConfig, err := config.File.GetSection("mysql")
	if err != nil {
		log.Println("mysql config error:", err)
		panic(err)
	}
	dst := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		mysqlConfig["user"],
		mysqlConfig["password"],
		mysqlConfig["host"],
		mysqlConfig["port"],
		mysqlConfig["dbname"],
		mysqlConfig["charset"],
	)
	Engine, err = xorm.NewEngine("mysql", dst)
	if err != nil {
		log.Println("mysql connect error:", err)
		panic(err)
	}

	err = Engine.Ping()
	if err != nil {
		log.Println("mysql Ping error:", err)
		panic(err)
	}

	maxIdle := config.File.MustInt("mysql", "max_idle", 2)
	maxConn := config.File.MustInt("mysql", "max_conn", 10)

	Engine.SetMaxIdleConns(maxIdle)
	Engine.SetMaxOpenConns(maxConn)
	Engine.ShowSQL(true)
	log.Println("mysql init success~")
}
