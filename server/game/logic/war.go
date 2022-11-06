package logic

import (
	"log"
	"ziweiSgServer/constant"
	"ziweiSgServer/db"
	"ziweiSgServer/server/common"
	"ziweiSgServer/server/game/model"
	"ziweiSgServer/server/game/model/data"
)

var WarService = &warService{}

type warService struct {
}

func (w *warService) GetWarReports(rid int) ([]model.WarReport, error) {
	mwr := make([]data.WarReport, 0)
	wr := &data.WarReport{}
	err := db.Engine.Table(wr).Where("a_rid=? or d_rid=?", rid, rid).Limit(30, 0).Desc("ctime").Find(&mwr)
	if err != nil {
		log.Println("GetWarReports查询战报错误", err)
		return nil, common.NewError(constant.DBError, "查询战报错误")
	}

	modelWarReports := make([]model.WarReport, 0)
	for _, v := range mwr {
		modelWarReports = append(modelWarReports, v.ToModel().(model.WarReport))
	}
	return modelWarReports, nil
}
