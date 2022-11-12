package logic

import (
	"encoding/json"
	"log"
	"sync"
	"xorm.io/xorm"
	"ziweiSgServer/constant"
	"ziweiSgServer/db"
	"ziweiSgServer/net"
	"ziweiSgServer/server/common"
	"ziweiSgServer/server/game/model"
	"ziweiSgServer/server/game/model/data"
)

var RoleAttributeService = &roleAttributeService{
	attrs: make(map[int]*data.RoleAttribute),
}

type roleAttributeService struct {
	mutex sync.RWMutex
	attrs map[int]*data.RoleAttribute
}

func (r *roleAttributeService) TryCreate(rid int, req *net.WsMsgReq) error {
	roleAttribute := &data.RoleAttribute{}
	ok, err := db.Engine.Table(roleAttribute).Where("rid=?", rid).Get(roleAttribute)
	if err != nil {
		log.Println("TryCreate查询角色属性出错", err)
		return common.NewError(constant.DBError, "数据库出错")
	}
	if ok {
		r.mutex.Lock()
		defer r.mutex.Unlock()
		r.attrs[rid] = roleAttribute
		return nil
	} else {
		// 如果查不到，则初始化
		roleAttribute.RId = rid
		roleAttribute.UnionId = 0
		roleAttribute.ParentId = 0
		roleAttribute.PosTags = ""
		session := req.Context.Get("dbSession").(*xorm.Session)
		if session != nil {
			_, err = session.Table(roleAttribute).Insert(roleAttribute)
		} else {
			_, err = db.Engine.Table(roleAttribute).Insert(roleAttribute)
		}
		if err != nil {
			log.Println("TryCreate插入角色属性出错", err)
			return common.NewError(constant.DBError, "数据库出错")
		}
		r.mutex.Lock()
		defer r.mutex.Unlock()
		r.attrs[rid] = roleAttribute
	}
	return nil
}

func (r *roleAttributeService) GetTagList(rid int) ([]model.PosTag, error) {
	r.mutex.RLock()
	roleAttribute, ok := r.attrs[rid]
	r.mutex.RUnlock()
	if !ok {
		roleAttribute := &data.RoleAttribute{}
		var err error
		ok, err = db.Engine.Table(roleAttribute).Where("rid=?", rid).Get(roleAttribute)
		if err != nil {
			log.Println("GetTagList查询标签列表出错", err)
			return nil, common.NewError(constant.DBError, "数据库错误")
		}
	}
	posTagList := make([]model.PosTag, 0)
	if ok {
		if roleAttribute.PosTags != "" {
			err := json.Unmarshal([]byte(roleAttribute.PosTags), &posTagList)
			if err != nil {
				return posTagList, common.NewError(constant.DBError, "数据库错误")
			}
		}
	}
	return posTagList, nil
}
