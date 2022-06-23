package login

import (
	"ziweiSgServer/db"
	"ziweiSgServer/net"
	"ziweiSgServer/server/login/controller"
)

var Router = net.NewRouter()

func Init() {
	// 测试数据库，并初始化mysql
	db.TestDB()
	// 还有别的初始化方法
	initRouter()
}

func initRouter() {
	controller.DefaultAccount.Router(Router)
}
