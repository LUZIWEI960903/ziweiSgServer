package logic

import (
	"encoding/json"
	"log"
	"ziweiSgServer/constant"
	"ziweiSgServer/db"
	"ziweiSgServer/server/common"
	"ziweiSgServer/server/game/gameConfig"
	"ziweiSgServer/server/game/model/data"
)

var CityFacilityService = &cityFacilityService{}

type cityFacilityService struct {
}

func (s cityFacilityService) TryCreate(cid, rid int) error {
	cf := &data.CityFacility{}
	ok, err := db.Engine.Table(cf).Where("cityId=?", cid).Get(cf)
	if err != nil {
		log.Println("TryCreate查询城市设施出错", err)
		return common.NewError(constant.DBError, "数据库错误")
	}
	if ok {
		return nil
	}
	cf.RId = rid
	cf.CityId = cid
	list := gameConfig.FacilityConf.List
	facs := make([]data.Facility, len(list))
	for index, v := range list {
		fac := data.Facility{
			Name:         v.Name,
			PrivateLevel: 0,
			Type:         v.Type,
			UpTime:       0,
		}
		facs[index] = fac
	}
	dataJson, _ := json.Marshal(facs)
	cf.Facilities = string(dataJson)
	_, err = db.Engine.Table(cf).Insert(cf)
	if err != nil {
		log.Println("TryCreate插入城市设施出错", err)
		return common.NewError(constant.DBError, "数据库错误")
	}
	return nil
}
