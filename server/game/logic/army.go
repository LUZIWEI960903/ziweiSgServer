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

func (s *armyService) GetArmysByCity(rid int, cityId int) ([]model.Army, error) {
	mrs := make([]data.Army, 0)
	mr := &data.Army{}
	err := db.Engine.Table(mr).Where("rid=? and cityId=?", rid, cityId).Find(&mrs)
	if err != nil {
		log.Println("GetArmysByCity军队查询出错", err)
		return nil, common.NewError(constant.DBError, "军队查询出错")
	}
	modelMrs := make([]model.Army, 0)
	for _, v := range mrs {
		modelMrs = append(modelMrs, v.ToModel().(model.Army))
	}
	return modelMrs, nil
}
