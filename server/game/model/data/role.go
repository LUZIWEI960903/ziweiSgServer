package data

import (
	"time"
	"ziweiSgServer/server/game/model"
)

type Role struct {
	RId        int       `xorm:"rid pk autoincr"`
	UId        int       `xorm:"uid"`
	NickName   string    `xorm:"nick_name" validate:"min=4,max=20,regexp=^[a-zA-Z0-9_]*$"`
	Balance    int       `xorm:"balance"`
	HeadId     int16     `xorm:"headId"`
	Sex        int8      `xorm:"sex"`
	Profile    string    `xorm:"profile"`
	LoginTime  time.Time `xorm:"login_time"`
	LogoutTime time.Time `xorm:"logout_time"`
	CreatedAt  time.Time `xorm:"created_at"`
}

func (r *Role) TableName() string {
	return "role"
}

func (r *Role) ToModel() interface{} {
	return model.Role{
		RId:      r.RId,
		UId:      r.UId,
		NickName: r.NickName,
		Sex:      r.Sex,
		Balance:  r.Balance,
		HeadId:   r.HeadId,
		Profile:  r.Profile,
	}
}
