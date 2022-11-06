package controller

import (
	"github.com/mitchellh/mapstructure"
	"ziweiSgServer/constant"
	"ziweiSgServer/net"
	"ziweiSgServer/server/common"
	"ziweiSgServer/server/game/logic"
	"ziweiSgServer/server/game/model"
	"ziweiSgServer/server/game/model/data"
)

var ArmyController = &armyController{}

type armyController struct {
}

func (a *armyController) Router(r *net.Router) {
	g := r.Group("army")
	g.AddRouter("myList", a.myList)
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
