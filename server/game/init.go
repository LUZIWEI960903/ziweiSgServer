package game

import (
	"ziweiSgServer/db"
	"ziweiSgServer/net"
	"ziweiSgServer/server/game/controller"
	"ziweiSgServer/server/game/gameConfig"
	"ziweiSgServer/server/game/gameConfig/general"
)

var Router = net.NewRouter()

func Init() {
	db.TestDB()
	// basic.json 加载基础配置
	gameConfig.Base.Load()
	// map_build.json 加载地图资源配置
	gameConfig.MapBuild.Load()
	// map.json 加载地图单元格配置
	gameConfig.MapRes.Load()
	// 加载城池设施配置
	gameConfig.FacilityConf.Load()
	// 加载武将配置信息
	general.General.Load()
	
	initRouter()
}

func initRouter() {
	controller.DefaultRoleController.Router(Router)
	controller.NationMapController.Router(Router)
	controller.GeneralController.Router(Router)
}
