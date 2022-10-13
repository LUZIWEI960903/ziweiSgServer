package logic

import (
	"log"
	"ziweiSgServer/constant"
	"ziweiSgServer/db"
	"ziweiSgServer/net"
	"ziweiSgServer/server/common"
	"ziweiSgServer/server/game/model/data"
)

var RoleAttributeService = &roleAttributeService{}

type roleAttributeService struct {
}

func (r *roleAttributeService) TryCreate(rid int, conn net.WSConn) error {
	roleAttribute := &data.RoleAttribute{}
	ok, err := db.Engine.Table(roleAttribute).Where("rid=?", rid).Get(roleAttribute)
	if err != nil {
		log.Println("TryCreate查询角色属性出错", err)
		return common.NewError(constant.DBError, "数据库出错")
	}
	if ok {
		return nil
	} else {
		// 如果查不到，则初始化
		roleAttribute.RId = rid
		roleAttribute.UnionId = 0
		roleAttribute.ParentId = 0
		_, err := db.Engine.Table(roleAttribute).Insert(roleAttribute)
		if err != nil {
			log.Println("TryCreate插入角色属性出错", err)
			return common.NewError(constant.DBError, "数据库出错")
		}
	}
	return nil
}
