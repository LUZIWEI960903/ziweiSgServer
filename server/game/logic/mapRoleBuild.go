package logic

import (
	"log"
	"ziweiSgServer/constant"
	"ziweiSgServer/db"
	"ziweiSgServer/server/common"
	"ziweiSgServer/server/game/model"
	"ziweiSgServer/server/game/model/data"
)

var MapRoleBuildService = &mapRoleBuildService{}

type mapRoleBuildService struct {
}

func (s *mapRoleBuildService) GetBuilds(rid int) ([]model.MapRoleBuild, error) {
	builds := make([]data.MapRoleBuild, 0)
	build := data.MapRoleBuild{}
	err := db.Engine.Table(build).Where("rid=?", rid).Find(&builds)
	modelBuilds := make([]model.MapRoleBuild, 0)
	if err != nil {
		log.Println("GetRoleBuilds查询玩家建筑错误", err)
		return modelBuilds, common.NewError(constant.DBError, "数据库错误")
	}
	for _, v := range builds {
		modelBuilds = append(modelBuilds, v.ToModel().(model.MapRoleBuild))
	}
	return modelBuilds, nil
}
