package controller

import (
	"ziweiSgServer/constant"
	"ziweiSgServer/net"
	"ziweiSgServer/server/game/gameConfig"
	"ziweiSgServer/server/game/middleware"
	"ziweiSgServer/server/game/model"
)

var NationMapController = &nationMapController{}

type nationMapController struct {
}

func (n *nationMapController) Router(r *net.Router) {
	g := r.Group("nationMap")
	g.Use(middleware.Log())
	g.AddRouter("config", n.config)
}

func (n *nationMapController) config(req *net.WsMsgReq, rsp *net.WsMsgRsp) {
	//reqObj := &model.ConfigReq{}
	rspObj := &model.ConfigRsp{}

	confs := gameConfig.MapBuild.Cfg
	rspObj.Confs = make([]model.Conf, len(confs))
	for index, v := range confs {
		rspObj.Confs[index].Type = v.Type
		rspObj.Confs[index].Grain = v.Grain
		rspObj.Confs[index].Stone = v.Stone
		rspObj.Confs[index].Level = v.Level
		rspObj.Confs[index].Iron = v.Iron
		rspObj.Confs[index].Wood = v.Wood
		rspObj.Confs[index].Durable = v.Durable
		rspObj.Confs[index].Name = v.Name
		rspObj.Confs[index].Defender = v.Defender
	}
	rsp.Body.Seq = req.Body.Seq
	rsp.Body.Name = req.Body.Name
	rsp.Body.Code = constant.OK
	rsp.Body.Msg = rspObj
}
