package data

import (
	"sync"
	"time"
	"ziweiSgServer/server/game/model"
)

// MapRoleCity 玩家城池
type MapRoleCity struct {
	mutex      sync.Mutex `xorm:"-"`
	CityId     int        `xorm:"cityId pk autoincr"`
	RId        int        `xorm:"rid"`
	Name       string     `xorm:"name" validate:"min=4,max=20,regexp=^[a-zA-Z0-9_]*$"`
	X          int        `xorm:"x"`
	Y          int        `xorm:"y"`
	IsMain     int8       `xorm:"is_main"`
	CurDurable int        `xorm:"cur_durable"`
	CreatedAt  time.Time  `xorm:"created_at"`
	OccupyTime time.Time  `xorm:"occupy_time"`
}

func (m *MapRoleCity) TableName() string {
	return "map_role_city"
}

func (m *MapRoleCity) ToModel() interface{} {
	return model.MapRoleCity{
		CityId:     m.CityId,
		RId:        m.RId,
		Name:       m.Name,
		UnionId:    0,
		UnionName:  "",
		ParentId:   0,
		X:          m.X,
		Y:          m.Y,
		IsMain:     m.IsMain == 1,
		Level:      1,
		CurDurable: m.CurDurable,
		MaxDurable: 1000,
		OccupyTime: m.OccupyTime.UnixNano() / 1e6,
	}
}
