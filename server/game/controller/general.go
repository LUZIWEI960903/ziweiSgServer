package controller

import (
	"ziweiSgServer/constant"
	"ziweiSgServer/net"
	"ziweiSgServer/server/common"
	"ziweiSgServer/server/game/logic"
	"ziweiSgServer/server/game/model"
	"ziweiSgServer/server/game/model/data"
)

var GeneralController = &generalController{}

type generalController struct {
}

func (c *generalController) Router(r *net.Router) {
	g := r.Group("general")
	g.AddRouter("myGenerals", c.myGenerals)
}

func (c *generalController) myGenerals(req *net.WsMsgReq, rsp *net.WsMsgRsp) {
	// 查询武将的时候  角色拥有的武将 查询出来即可
	// 如果初始化 进入游戏 武将没有 需要随机三个武将
	//reqObj := &model.MyGeneralReq{}
	rspObj := &model.MyGeneralRsp{}

	role, err := req.Conn.GetProperty("role")
	if err != nil {
		rsp.Body.Code = constant.SessionInvalid
		return
	}

	rid := role.(*data.Role).RId

	gs, err := logic.GeneralService.GetGenerals(rid)
	if err != nil {
		rsp.Body.Code = err.(*common.MyError).Code()
		return
	}
	rspObj.Generals = gs

	rsp.Body.Seq = req.Body.Seq
	rsp.Body.Name = req.Body.Name
	rsp.Body.Code = constant.OK
	rsp.Body.Msg = rspObj
}
