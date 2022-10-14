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
	// map_build.json 加载地图资源配置
	gameConfig.MapBuild.Load()
	// map.json 加载地图单元格配置
	gameConfig.MapRes.Load()

	db.TestDB()
	initRouter()
}

func initRouter() {
	controller.DefaultRoleController.Router(Router)
	controller.NationMapController.Router(Router)
}
