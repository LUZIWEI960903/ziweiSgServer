package logic

import (
	"log"
	"ziweiSgServer/constant"
	"ziweiSgServer/db"
	"ziweiSgServer/server/common"
	"ziweiSgServer/server/game/model"
	"ziweiSgServer/server/game/model/data"
)

var ArmyService = &armyService{}

type armyService struct {
}

func (s *armyService) GetArmys(rid int) ([]model.Army, error) {
	armys := make([]data.Army, 0)
	army := &data.Army{}
	err := db.Engine.Table(army).Where("rid=?", rid).Find(&armys)
	modelAmrys := make([]model.Army, 0)
	if err != nil {
		log.Println("GetArmys获取军队错误", err)
		return modelAmrys, common.NewError(constant.DBError, "数据库错误")
	}
	for _, v := range armys {
		modelAmrys = append(modelAmrys, v.ToModel().(model.Army))
	}
	return modelAmrys, nil
}
