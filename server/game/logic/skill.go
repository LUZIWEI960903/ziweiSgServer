package logic

import (
	"log"
	"ziweiSgServer/constant"
	"ziweiSgServer/db"
	"ziweiSgServer/server/common"
	"ziweiSgServer/server/game/model"
	"ziweiSgServer/server/game/model/data"
)

var SkillService = &skillService{}

type skillService struct {
}

func (s *skillService) GetSkillList(rid int) ([]model.Skill, error) {
	dsl := make([]data.Skill, 0)
	skill := &data.Skill{}
	err := db.Engine.Table(skill).Where("rid=?", rid).Find(&dsl)
	if err != nil {
		log.Println("GetSkillList查询技能列表错误", err)
		return nil, common.NewError(constant.DBError, "技能查询出错")
	}
	modelSL := make([]model.Skill, 0)
	for _, v := range dsl {
		modelSL = append(modelSL, v.ToModel().(model.Skill))
	}
	return modelSL, nil
}
