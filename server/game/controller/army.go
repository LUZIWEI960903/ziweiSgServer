package controller

import (
	"ziweiSgServer/constant"
	"ziweiSgServer/net"
	"ziweiSgServer/server/common"
	"ziweiSgServer/server/game/logic"
	"ziweiSgServer/server/game/middleware"
	"ziweiSgServer/server/game/model"
	"ziweiSgServer/server/game/model/data"

	"github.com/mitchellh/mapstructure"
)

var ArmyController = &armyController{}

type armyController struct {
}

func (a *armyController) Router(r *net.Router) {
	g := r.Group("army")
	g.Use(middleware.Log())
	g.AddRouter("myList", a.myList, middleware.CheckRole())
}

func (a *armyController) myList(req *net.WsMsgReq, rsp *net.WsMsgRsp) {
	reqObj := &model.ArmyListReq{}
	rspObj := &model.ArmyListRsp{}

	err := mapstructure.Decode(req.Body.Msg, reqObj)
	if err != nil {
		rsp.Body.Code = constant.InvalidParam
		return
	}

	role, err := req.Conn.GetProperty("role")
	if err != nil {
		rsp.Body.Code = constant.SessionInvalid
		return
	}

	rid := role.(*data.Role).RId
	arms, err := logic.ArmyService.GetArmysByCity(rid, reqObj.CityId)
	if err != nil {
		rsp.Body.Code = err.(*common.MyError).Code()
		return
	}
	rspObj.Armys = arms
	rspObj.CityId = reqObj.CityId

	rsp.Body.Seq = req.Body.Seq
	rsp.Body.Name = req.Body.Name
	rsp.Body.Code = constant.OK
	rsp.Body.Msg = rspObj
}
