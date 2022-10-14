package data

import (
	"time"
	"ziweiSgServer/server/game/model"
)

const (
	GeneralNormal      = 0 //正常
	GeneralComposeStar = 1 //星级合成
	GeneralConvert     = 2 //转换
)
const SkillLimit = 3

type General struct {
	Id            int             `xorm:"id pk autoincr"`
	RId           int             `xorm:"rid"`
	CfgId         int             `xorm:"cfgId"`
	PhysicalPower int             `xorm:"physical_power"`
	Level         int8            `xorm:"level"`
	Exp           int             `xorm:"exp"`
	Order         int8            `xorm:"order"`
	CityId        int             `xorm:"cityId"`
	CreatedAt     time.Time       `xorm:"created_at"`
	CurArms       int             `xorm:"arms"`
	HasPrPoint    int             `xorm:"has_pr_point"`
	UsePrPoint    int             `xorm:"use_pr_point"`
	AttackDis     int             `xorm:"attack_distance"`
	ForceAdded    int             `xorm:"force_added"`
	StrategyAdded int             `xorm:"strategy_added"`
	DefenseAdded  int             `xorm:"defense_added"`
	SpeedAdded    int             `xorm:"speed_added"`
	DestroyAdded  int             `xorm:"destroy_added"`
	StarLv        int8            `xorm:"star_lv"`
	Star          int8            `xorm:"star"`
	ParentId      int             `xorm:"parentId"`
	Skills        string          `xorm:"skills"`
	SkillsArray   []*model.GSkill `xorm:"-"`
	State         int8            `xorm:"state"`
}

func (g *General) TableName() string {
	return "general"
}

func (g *General) ToModel() interface{} {
	return model.General{
		Id:            g.Id,
		CfgId:         g.CfgId,
		PhysicalPower: g.PhysicalPower,
		Order:         g.Order,
		Level:         g.Level,
		Exp:           g.Exp,
		CityId:        g.CityId,
		CurArms:       g.CurArms,
		HasPrPoint:    g.HasPrPoint,
		UsePrPoint:    g.UsePrPoint,
		AttackDis:     g.AttackDis,
		ForceAdded:    g.ForceAdded,
		StrategyAdded: g.StrategyAdded,
		DefenseAdded:  g.DefenseAdded,
		SpeedAdded:    g.SpeedAdded,
		DestroyAdded:  g.DestroyAdded,
		StarLv:        g.StarLv,
		Star:          g.Star,
		ParentId:      g.ParentId,
		Skills:        g.SkillsArray,
		State:         g.State,
	}
}
