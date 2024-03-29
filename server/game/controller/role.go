package controller

import (
	"ziweiSgServer/constant"
	"ziweiSgServer/net"
	"ziweiSgServer/server/common"
	"ziweiSgServer/server/game/logic"
	"ziweiSgServer/server/game/middleware"
	"ziweiSgServer/server/game/model"
	"ziweiSgServer/server/game/model/data"
	"ziweiSgServer/utils"

	"github.com/mitchellh/mapstructure"
)

var DefaultRoleController = &RoleController{}

type RoleController struct {
}

func (r *RoleController) Router(router *net.Router) {
	g := router.Group("role")
	g.Use(middleware.Log())
	g.AddRouter("enterServer", r.enterServer)
	g.AddRouter("myProperty", r.myProperty, middleware.CheckRole())
	g.AddRouter("posTagList", r.posTagList, middleware.CheckRole())
}

func (r *RoleController) enterServer(req *net.WsMsgReq, rsp *net.WsMsgRsp) {
	// 进入游戏的逻辑
	// session 需要验证是否合法， 合法 取出登录用户的 id
	// 根据用户 id 查询对应的游戏角色，如果有 就继续，没有 提示无角色
	// 根据角色id 查询角色拥有的资源 roleRes，如果有则返回，没有 初始化资源
	reqObj := &model.EnterServerReq{}
	rspObj := &model.EnterServerRsp{}

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
	// 初始化相关的逻辑
	err = logic.RoleService.EnterServer(uid, rspObj, req)
	if err != nil {
		rsp.Body.Code = err.(*common.MyError).Code()
		return
	}
	rsp.Body.Seq = req.Body.Seq
	rsp.Body.Name = req.Body.Name
	rsp.Body.Code = constant.OK
	rsp.Body.Msg = rspObj
}

func (r *RoleController) myProperty(req *net.WsMsgReq, rsp *net.WsMsgRsp) {
	// 分别根据角色id 去查询 军队 资源 建筑 城池 武将
	//reqObj := &model.MyRolePropertyReq{}
	rspObj := &model.MyRolePropertyRsp{}

	role, err := req.Conn.GetProperty("role")
	if err != nil {
		rsp.Body.Code = constant.SessionInvalid
		return
	}
	rid := role.(*data.Role).RId

	// 资源
	rspObj.RoleRes, err = logic.RoleService.GetRoleRes(rid)
	if err != nil {
		rsp.Body.Code = err.(*common.MyError).Code()
		return
	}
	// 城池
	rspObj.Citys, err = logic.MapRoleCityService.GetRoleCitys(rid)
	if err != nil {
		rsp.Body.Code = err.(*common.MyError).Code()
		return
	}
	// 建筑
	rspObj.MRBuilds, err = logic.MapRoleBuildService.GetBuilds(rid)
	if err != nil {
		rsp.Body.Code = err.(*common.MyError).Code()
		return
	}
	// 军队
	rspObj.Armys, err = logic.ArmyService.GetArmys(rid)
	if err != nil {
		rsp.Body.Code = err.(*common.MyError).Code()
		return
	}
	// 武将
	rspObj.Generals, err = logic.GeneralService.GetGenerals(rid)
	if err != nil {
		rsp.Body.Code = err.(*common.MyError).Code()
		return
	}

	rsp.Body.Seq = req.Body.Seq
	rsp.Body.Name = req.Body.Name
	rsp.Body.Code = constant.OK
	rsp.Body.Msg = rspObj
}

func (r *RoleController) posTagList(req *net.WsMsgReq, rsp *net.WsMsgRsp) {
	//reqObj := &model.PosTagListReq{}
	rspObj := &model.PosTagListRsp{}

	role, err := req.Conn.GetProperty("role")
	if err != nil {
		rsp.Body.Code = constant.SessionInvalid
		return
	}
	rid := role.(*data.Role).RId
	rspObj.PosTags, err = logic.RoleAttributeService.GetTagList(rid)
	if err != nil {
		rsp.Body.Code = err.(*common.MyError).Code()
		return
	}
	rsp.Body.Seq = req.Body.Seq
	rsp.Body.Name = req.Body.Name
	rsp.Body.Code = constant.OK
	rsp.Body.Msg = rspObj
}
