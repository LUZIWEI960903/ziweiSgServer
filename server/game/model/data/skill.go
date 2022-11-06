package data

import "ziweiSgServer/server/game/model"

//军队
type Skill struct {
	Id             int    `xorm:"id pk autoincr"`
	RId            int    `xorm:"rid"`
	CfgId          int    `xorm:"cfgId"`
	BelongGenerals string `xorm:"belong_generals"`
	Generals       []int  `xorm:"-"`
}

func NewSkill(rid int, cfgId int) *Skill {
	return &Skill{
		CfgId:          cfgId,
		RId:            rid,
		Generals:       []int{},
		BelongGenerals: "[]",
	}
}

func (s *Skill) TableName() string {
	return "skill"
}

func (s *Skill) ToModel() interface{} {
	return model.Skill{
		Id:       s.Id,
		CfgId:    s.CfgId,
		Generals: s.Generals,
	}
}
