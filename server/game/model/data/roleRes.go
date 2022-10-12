package data

import "ziweiSgServer/server/game/model"

type RoleRes struct {
	Id     int `xorm:"id pk autoincr"`
	RId    int `xorm:"rid"`
	Wood   int `xorm:"wood"`
	Iron   int `xorm:"iron"`
	Stone  int `xorm:"stone"`
	Grain  int `xorm:"grain"`
	Gold   int `xorm:"gold"`
	Decree int `xorm:"decree"` //令牌
}

func (r *RoleRes) TableName() string {
	return "role_res"
}

func (r *RoleRes) ToModel() interface{} {
	return model.RoleRes{
		Wood:          r.Wood,
		Iron:          r.Iron,
		Stone:         r.Stone,
		Grain:         r.Grain,
		Gold:          r.Gold,
		Decree:        r.Decree,
		WoodYield:     1,
		IronYield:     1,
		StoneYield:    1,
		GrainYield:    1,
		GoldYield:     1,
		DepotCapacity: 10000,
	}
}
