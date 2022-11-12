package logic

import (
	"log"
	"math/rand"
	"time"
	"xorm.io/xorm"
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

func (m *mapRoleCityService) InitCity(rid int, roleNickName string, req *net.WsMsgReq) error {
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
		for {
			mapRoleCity.X = rand.Intn(global.MapWidth)
			mapRoleCity.Y = rand.Intn(global.MapHeight)

			// 这个城池 能不能在这个坐标点创建 需要判断 系统城池五格之内 不能有玩家的城池
			if IsCanBuild(mapRoleCity.X, mapRoleCity.Y) {
				mapRoleCity.RId = rid
				mapRoleCity.Name = roleNickName
				mapRoleCity.IsMain = 1
				mapRoleCity.CurDurable = gameConfig.Base.City.Durable
				mapRoleCity.CreatedAt = time.Now()
				session := req.Context.Get("dbSession").(*xorm.Session)
				if session != nil {
					_, err = session.Table(mapRoleCity).Insert(mapRoleCity)
				} else {
					_, err = db.Engine.Table(mapRoleCity).Insert(mapRoleCity)
				}
				if err != nil {
					log.Println("InitCity插入角色城池出错", err)
					return common.NewError(constant.DBError, "数据库出错")
				}
				// 初始化城池的设施
				if err := CityFacilityService.TryCreate(mapRoleCity.CityId, rid, req); err != nil {
					log.Println("InitCity插入城池设施出错", err)
					return common.NewError(err.(*common.MyError).Code(), err.Error())
				}
				break
			}
		}
	}
	return nil
}

func IsCanBuild(x int, y int) bool {
	confs := gameConfig.MapRes.Confs
	pIndex := global.ToPosition(x, y)
	_, ok := confs[pIndex]
	if !ok {
		return false
	}
	sysBuild := gameConfig.MapRes.SysBuild
	// 系统城池的5格内 不能创建玩家城池
	for _, v := range sysBuild {
		if v.Type == gameConfig.MapBuildSysCity {
			if x <= v.X+5 &&
				x >= v.X-5 &&
				y <= v.Y+5 &&
				y >= v.Y-5 {
				return false
			}
		}
	}
	return true
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
