package data

import (
	"time"
	"ziweiSgServer/server/game/model"
)

type WarReport struct {
	Id                int       `xorm:"id pk autoincr"`
	AttackRid         int       `xorm:"a_rid"`
	DefenseRid        int       `xorm:"d_rid"`
	BegAttackArmy     string    `xorm:"b_a_army"`
	BegDefenseArmy    string    `xorm:"b_d_army"`
	EndAttackArmy     string    `xorm:"e_a_army"`
	EndDefenseArmy    string    `xorm:"e_d_army"`
	BegAttackGeneral  string    `xorm:"b_a_general"`
	BegDefenseGeneral string    `xorm:"b_d_general"`
	EndAttackGeneral  string    `xorm:"e_a_general"`
	EndDefenseGeneral string    `xorm:"e_d_general"`
	Result            int       `xorm:"result"` //0失败，1打平，2胜利
	Rounds            string    `xorm:"rounds"` //回合
	AttackIsRead      bool      `xorm:"a_is_read"`
	DefenseIsRead     bool      `xorm:"d_is_read"`
	DestroyDurable    int       `xorm:"destroy"`
	Occupy            int       `xorm:"occupy"`
	X                 int       `xorm:"x"`
	Y                 int       `xorm:"y"`
	CTime             time.Time `xorm:"ctime"`
}

func (w *WarReport) TableName() string {
	return "war_report"
}

func (w *WarReport) ToModel() interface{} {
	return model.WarReport{
		Id:                w.Id,
		AttackRid:         w.AttackRid,
		DefenseRid:        w.DefenseRid,
		BegAttackArmy:     w.BegAttackArmy,
		BegDefenseArmy:    w.BegDefenseArmy,
		EndAttackArmy:     w.EndAttackArmy,
		EndDefenseArmy:    w.EndDefenseArmy,
		BegAttackGeneral:  w.BegAttackGeneral,
		BegDefenseGeneral: w.BegDefenseGeneral,
		EndAttackGeneral:  w.EndAttackGeneral,
		EndDefenseGeneral: w.EndDefenseGeneral,
		Result:            w.Result,
		Rounds:            w.Rounds,
		AttackIsRead:      w.AttackIsRead,
		DefenseIsRead:     w.DefenseIsRead,
		DestroyDurable:    w.DestroyDurable,
		Occupy:            w.Occupy,
		X:                 w.X,
		Y:                 w.Y,
		CTime:             int(w.CTime.UnixNano() / 1e6),
	}
}
