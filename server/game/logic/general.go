package logic

import (
	"log"
	"ziweiSgServer/constant"
	"ziweiSgServer/db"
	"ziweiSgServer/server/common"
	"ziweiSgServer/server/game/model"
	"ziweiSgServer/server/game/model/data"
)

var GeneralService = &generalService{}

type generalService struct {
}

func (s *generalService) GetGenerals(rid int) ([]model.General, error) {
	generals := make([]data.General, 0)
	general := &data.General{}
	err := db.Engine.Table(general).Where("rid=?", rid).Find(&generals)
	modelGenerals := make([]model.General, 0)
	if err != nil {
		log.Println("GetGenerals获取武将错误", err)
		return modelGenerals, common.NewError(constant.DBError, "数据库错误")
	}
	for _, v := range generals {
		modelGenerals = append(modelGenerals, v.ToModel().(model.General))
	}
	return modelGenerals, nil
}
