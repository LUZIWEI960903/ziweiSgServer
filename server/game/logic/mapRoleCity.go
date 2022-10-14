package logic

import (
	"log"
	"math/rand"
	"time"
	"ziweiSgServer/constant"
	"ziweiSgServer/db"
	"ziweiSgServer/net"
	"ziweiSgServer/server/common"
	"ziweiSgServer/server/game/gameConfig"
	"ziweiSgServer/server/game/global"
	"ziweiSgServer/server/game/model"
	"ziweiSgServer/server/game/model/data"
)

var MapRoleCityService = &mapRoleCityService{}

type mapRoleCityService struct {
}

func (m *mapRoleCityService) InitCity(rid int, roleNickName string, conn net.WSConn) error {
	mapRoleCity := &data.MapRoleCity{}
	ok, err := db.Engine.Table(mapRoleCity).Where("rid=?", rid).Get(mapRoleCity)
	if err != nil {
		log.Println("InitCity查询角色城池出错", err)
		return common.NewError(constant.DBError, "数据库出错")
	}
	if ok {
		return nil
	} else {
		// 如果查不到，则初始化
		mapRoleCity.X = rand.Intn(global.MapWidth)
		mapRoleCity.Y = rand.Intn(global.MapHeight)
		// 这个城池 能不能在这个坐标点创建 需要判断 五格之内 不能有玩家的城池
		// TODO
		mapRoleCity.RId = rid
		mapRoleCity.Name = roleNickName
		mapRoleCity.IsMain = 1
		mapRoleCity.CurDurable = gameConfig.Base.City.Durable
		mapRoleCity.CreatedAt = time.Now()

		_, err := db.Engine.Table(mapRoleCity).Insert(mapRoleCity)
		if err != nil {
			log.Println("InitCity插入角色城池出错", err)
			return common.NewError(constant.DBError, "数据库出错")
		}
		// 初始化城池的设施
		// TODO
	}
	return nil
}

func (m *mapRoleCityService) GetRoleCitys(rid int) ([]model.MapRoleCity, error) {
	citys := make([]data.MapRoleCity, 0)
	city := &data.MapRoleCity{}
	err := db.Engine.Table(city).Where("rid=?", rid).Find(&citys)
	modelCitys := make([]model.MapRoleCity, 0)
	if err != nil {
		log.Println("GetRoleCitys获取角色城池出错", err)
		return modelCitys, common.NewError(constant.DBError, "数据库出错")
	}
	for _, v := range citys {
		modelCitys = append(modelCitys, v.ToModel().(model.MapRoleCity))
	}
	return modelCitys, nil
}
