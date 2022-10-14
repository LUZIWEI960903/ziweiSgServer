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
	// map_build.json 加载地图基础配置
	gameConfig.MapBuild.Load()
	db.TestDB()
	initRouter()
}

func initRouter() {
	controller.DefaultRoleController.Router(Router)
	controller.NationMapController.Router(Router)
}
