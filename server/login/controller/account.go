package controller

import (
	"ziweiSgServer/net"
	"ziweiSgServer/server/login/proto"
)

var DefaultAccount = &Account{}

type Account struct {
}

func (a *Account) Router(r *net.Router) {
	g := r.Group("account")
	g.AddRouter("login", a.login)
}

func (a *Account) login(req *net.WsMsgReq, rsp *net.WsMsgRsp) {
	rsp.Body.Code = 0
	loginRsp := &proto.LoginRsp{}
	loginRsp.UId = 1
	loginRsp.Username = "admin"
	loginRsp.Session = "as"
	loginRsp.Password = ""
	rsp.Body.Msg = loginRsp
}
