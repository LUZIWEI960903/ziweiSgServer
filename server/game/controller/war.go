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

var WarController = &warController{}

type warController struct {
}

func (w *warController) Router(r *net.Router) {
	g := r.Group("war")
	g.Use(middleware.Log())
	g.AddRouter("report", w.report, middleware.CheckRole())
}

func (w *warController) report(req *net.WsMsgReq, rsp *net.WsMsgRsp) {
	//reqObj := &model.WarReportReq{}
	rspObj := &model.WarReportRsp{}

	role, err := req.Conn.GetProperty("role")
	if err != nil {
		rsp.Body.Code = constant.SessionInvalid
		return
	}

	rid := role.(*data.Role).RId

	wrs, err := logic.WarService.GetWarReports(rid)
	if err != nil {
		rsp.Body.Code = err.(*common.MyError).Code()
		return
	}
	rspObj.List = wrs

	rsp.Body.Seq = req.Body.Seq
	rsp.Body.Name = req.Body.Name
	rsp.Body.Code = constant.OK
	rsp.Body.Msg = rspObj
}
