package data

import (
	"time"
	"ziweiSgServer/server/game/model"
)

const (
	MapBuildSysFortress = 50 //系统要塞
	MapBuildSysCity     = 51 //系统城市
	MapBuildFortress    = 56 //玩家要塞
)

type MapRoleBuild struct {
	Id         int       `xorm:"id pk autoincr"`
	RId        int       `xorm:"rid"`
	Type       int8      `xorm:"type"`
	Level      int8      `xorm:"level"`
	OPLevel    int8      `xorm:"op_level"` //操作level
	X          int       `xorm:"x"`
	Y          int       `xorm:"y"`
	Name       string    `xorm:"name"`
	Wood       int       `xorm:"-"`
	Iron       int       `xorm:"-"`
	Stone      int       `xorm:"-"`
	Grain      int       `xorm:"-"`
	Defender   int       `xorm:"-"`
	CurDurable int       `xorm:"cur_durable"`
	MaxDurable int       `xorm:"max_durable"`
	OccupyTime time.Time `xorm:"occupy_time"`
	EndTime    time.Time `xorm:"end_time"` //建造或升级完的时间
	GiveUpTime int64     `xorm:"giveUp_time"`
}

func (m *MapRoleBuild) TableName() string {
	return "map_role_build"
}

func (m *MapRoleBuild) ToModel() interface{} {
	return model.MapRoleBuild{
		RId:        m.RId,
		RNick:      "111",
		Name:       m.Name,
		UnionId:    0,
		UnionName:  "",
		ParentId:   0,
		X:          m.X,
		Y:          m.Y,
		Type:       m.Type,
		Level:      m.Level,
		OPLevel:    m.OPLevel,
		CurDurable: m.CurDurable,
		MaxDurable: m.MaxDurable,
		Defender:   m.Defender,
		OccupyTime: m.OccupyTime.UnixNano() / 1e6,
		EndTime:    m.EndTime.UnixNano() / 1e6,
		GiveUpTime: m.GiveUpTime * 1000,
	}
}
