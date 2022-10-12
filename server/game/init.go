package game

import (
	"ziweiSgServer/db"
	"ziweiSgServer/net"
	"ziweiSgServer/server/game/controller"
	"ziweiSgServer/server/game/gameConfig"
)

var Router = net.NewRouter()

func Init() {
	// basic.json 加载基础配置
	gameConfig.Base.Load()
	db.TestDB()
	initRouter()
}

func initRouter() {
	controller.DefaultRoleController.Router(Router)
}
