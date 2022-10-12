package controller

import (
	"github.com/mitchellh/mapstructure"
	"log"
	"time"
	"ziweiSgServer/constant"
	"ziweiSgServer/db"
	"ziweiSgServer/net"
	"ziweiSgServer/server/game/gameConfig"
	"ziweiSgServer/server/game/model"
	"ziweiSgServer/server/game/model/data"
	"ziweiSgServer/utils"
)

var DefaultRoleController = &RoleController{}

type RoleController struct {
}

func (r *RoleController) Router(router *net.Router) {
	g := router.Group("role")
	g.AddRouter("enterServer", r.enterServer)
}

func (r *RoleController) enterServer(req *net.WsMsgReq, rsp *net.WsMsgRsp) {
	// 进入游戏的逻辑
	// session 需要验证是否合法， 合法 取出登录用户的 id
	// 根据用户 id 查询对应的游戏角色，如果有 就继续，没有 提示无角色
	// 根据角色id 查询角色拥有的资源 roleRes，如果有则返回，没有 初始化资源
	reqObj := &model.EnterServerReq{}
	rspObj := &model.EnterServerRsp{}
	rsp.Body.Seq = req.Body.Seq
	rsp.Body.Name = req.Body.Name

	err := mapstructure.Decode(req.Body.Msg, reqObj)
	if err != nil {
		rsp.Body.Code = constant.InvalidParam
		return
	}
	session := reqObj.Session
	_, claims, err := utils.ParseToken(session)
	if err != nil {
		rsp.Body.Code = constant.SessionInvalid
		return
	}
	uid := claims.Uid
	role := &data.Role{}
	ok, err := db.Engine.Table(role).Where("uid=?", uid).Get(role)
	if err != nil {
		log.Println("enterServer查询角色出错", err)
		rsp.Body.Code = constant.DBError
		return
	}
	if ok {
		rsp.Body.Code = constant.OK
		rsp.Body.Msg = rspObj

		rid := role.RId
		// 查询角色对应的 资源
		roleRes := &data.RoleRes{}
		ok, err := db.Engine.Table(roleRes).Where("rid=?", rid).Get(roleRes)
		if err != nil {
			log.Println("enterServer查询角色资源出错", err)
			rsp.Body.Code = constant.DBError
			return
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
				rsp.Body.Code = constant.DBError
				return
			}
		}
		rspObj.RoleRes = roleRes.ToModel().(model.RoleRes)
		rspObj.Role = role.ToModel().(model.Role)
		rspObj.Time = time.Now().UnixNano() / 1e6
		token, _ := utils.Award(rid)
		rspObj.Token = token
	} else {
		rsp.Body.Code = constant.RoleNotExist
		return
	}
}
