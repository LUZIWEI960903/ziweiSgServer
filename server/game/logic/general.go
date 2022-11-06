package logic

import (
	"encoding/json"
	"log"
	"time"
	"ziweiSgServer/constant"
	"ziweiSgServer/db"
	"ziweiSgServer/server/common"
	"ziweiSgServer/server/game/gameConfig"
	"ziweiSgServer/server/game/gameConfig/general"
	"ziweiSgServer/server/game/model"
	"ziweiSgServer/server/game/model/data"
)

var GeneralService = &generalService{}

type generalService struct {
}

func (s *generalService) GetGenerals(rid int) ([]model.General, error) {
	generals := make([]*data.General, 0)
	g := &data.General{}
	err := db.Engine.Table(g).Where("rid=?", rid).Find(&generals)
	if err != nil {
		log.Println("GetGenerals获取武将错误", err)
		return nil, common.NewError(constant.DBError, "数据库错误")
	}
	if len(generals) <= 0 {
		//没有武将，随机3个武将
		for i := 0; i < 3; i++ {
			cfgId := general.General.Rand()
			gen, err := s.NewGeneral(cfgId, rid, 0)
			if err != nil {
				log.Println(err)
				continue
			}

			generals = append(generals, gen)
		}

	}
	modelGenerals := make([]model.General, 0)
	for _, v := range generals {
		modelGenerals = append(modelGenerals, v.ToModel().(model.General))
	}
	return modelGenerals, nil
}

const (
	GeneralNormal      = 0 //正常
	GeneralComposeStar = 1 //星级合成
	GeneralConvert     = 2 //转换
)

func (s *generalService) NewGeneral(cfgId int, rid int, level int8) (*data.General, error) {
	cfg := general.General.GMap[cfgId]
	//初始 武将无技能 但是有三个技能槽
	sa := make([]*model.GSkill, 3)
	ss, _ := json.Marshal(sa)
	gen := &data.General{
		PhysicalPower: gameConfig.Base.General.PhysicalPowerLimit,
		RId:           rid,
		CfgId:         cfg.CfgId,
		Order:         0,
		CityId:        0,
		Level:         level,
		CreatedAt:     time.Now(),
		CurArms:       cfg.Arms[0],
		HasPrPoint:    0,
		UsePrPoint:    0,
		AttackDis:     0,
		ForceAdded:    0,
		StrategyAdded: 0,
		DefenseAdded:  0,
		SpeedAdded:    0,
		DestroyAdded:  0,
		Star:          cfg.Star,
		StarLv:        0,
		ParentId:      0,
		SkillsArray:   sa,
		Skills:        string(ss),
		State:         GeneralNormal,
	}

	_, err := db.Engine.Table(gen).Insert(gen)
	if err != nil {
		log.Println("GetGenerals插入武将错误", err)
		return nil, err
	}
	return gen, nil
}
