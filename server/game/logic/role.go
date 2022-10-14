package logic

import (
	"log"
	"time"
	"ziweiSgServer/constant"
	"ziweiSgServer/db"
	"ziweiSgServer/net"
	"ziweiSgServer/server/common"
	"ziweiSgServer/server/game/gameConfig"
	"ziweiSgServer/server/game/model"
	"ziweiSgServer/server/game/model/data"
	"ziweiSgServer/utils"
)

var RoleService = &roleService{}

type roleService struct {
}

func (r *roleService) EnterServer(uid int, rspObj *model.EnterServerRsp, conn net.WSConn) error {
	role := &data.Role{}
	ok, err := db.Engine.Table(role).Where("uid=?", uid).Get(role)
	if err != nil {
		log.Println("enterServer查询角色出错", err)
		return common.NewError(constant.DBError, "数据库出错")
	}
	if ok {
		rid := role.RId
		// 查询角色对应的 资源
		roleRes := &data.RoleRes{}
		ok, err := db.Engine.Table(roleRes).Where("rid=?", rid).Get(roleRes)
		if err != nil {
			log.Println("enterServer查询角色资源出错", err)
			return common.NewError(constant.DBError, "数据库出错")
		}
		if !ok {
			roleRes.RId = rid
			roleRes.Wood = gameConfig.Base.Role.Wood
			roleRes.Iron = gameConfig.Base.Role.Iron
			roleRes.Stone = gameConfig.Base.Role.Stone
			roleRes.Grain = gameConfig.Base.Role.Grain
			roleRes.Gold = gameConfig.Base.Role.Gold
			roleRes.Decree = gameConfig.Base.Role.Decree
			_, err := db.Engine.Table(roleRes).Insert(roleRes)
			if err != nil {
				log.Println("enterServer插入角色资源出错", err)
				return common.NewError(constant.DBError, "数据库出错")
			}
		}
		rspObj.RoleRes = roleRes.ToModel().(model.RoleRes)
		rspObj.Role = role.ToModel().(model.Role)
		rspObj.Time = time.Now().UnixNano() / 1e6
		token, _ := utils.Award(rid)
		rspObj.Token = token

		conn.SetProperty("role", role)

		// 初始化玩家属性
		if err := RoleAttributeService.TryCreate(rid, conn); err != nil {
			return common.NewError(constant.DBError, "数据库错误")
		}

		// 初始化城池
		if err := MapRoleCityService.InitCity(rid, role.NickName, conn); err != nil {
			return common.NewError(constant.DBError, "数据库错误")
		}

	} else {
		return common.NewError(constant.RoleNotExist, "角色不存在")
	}
	return nil
}

func (r *roleService) GetRoleRes(rid int) (model.RoleRes, error) {
	roleRes := &data.RoleRes{}
	ok, err := db.Engine.Table(roleRes).Where("rid=?", rid).Get(roleRes)
	if err != nil {
		log.Println("MyProperty角色资源出错", err)
		return model.RoleRes{}, common.NewError(constant.DBError, "数据库错误")
	}
	if !ok {
		log.Println("MyProperty角色资源出错", err)
		return model.RoleRes{}, common.NewError(constant.DBError, "角色资源不存在")
	}
	return roleRes.ToModel().(model.RoleRes), nil
}
