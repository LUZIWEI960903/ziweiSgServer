package controller

import (
	"ziweiSgServer/constant"
	"ziweiSgServer/net"
	"ziweiSgServer/server/common"
	"ziweiSgServer/server/game/logic"
	"ziweiSgServer/server/game/middleware"
	"ziweiSgServer/server/game/model"
	"ziweiSgServer/server/game/model/data"
)

var SkillController = &skillController{}

type skillController struct {
}

func (s *skillController) Router(r *net.Router) {
	g := r.Group("skill")
	g.Use(middleware.Log())
	g.AddRouter("list", s.list, middleware.CheckRole())
}

func (s *skillController) list(req *net.WsMsgReq, rsp *net.WsMsgRsp) {
	//reqObj := &model.SkillListReq{}
	rspObj := &model.SkillListRsp{}

	role, err := req.Conn.GetProperty("role")
	if err != nil {
		rsp.Body.Code = constant.SessionInvalid
		return
	}

	rid := role.(*data.Role).RId
	sls, err := logic.SkillService.GetSkillList(rid)
	if err != nil {
		rsp.Body.Code = err.(*common.MyError).Code()
		return
	}
	rspObj.List = sls

	rsp.Body.Seq = req.Body.Seq
	rsp.Body.Name = req.Body.Name
	rsp.Body.Code = constant.OK
	rsp.Body.Msg = rspObj
}
