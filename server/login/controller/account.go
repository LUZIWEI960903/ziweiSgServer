package controller

import (
	"github.com/mitchellh/mapstructure"
	"log"
	"ziweiSgServer/db"
	"ziweiSgServer/net"
	"ziweiSgServer/server/login/model"
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
	/*
		1. req中获取 用户名 密码 硬件id
		2. 根据用户名查询user表 得到数据
		3. 密码校验
		4. 保存用户登录记录
		5. 保存用户最后一次登录的信息
		6. 客户端需要一个session，jwt
		7. 客户端在发起需要用户登录的行为时，判断用户是否合法
	*/

	loginReq := &proto.LoginReq{}
	loginRsp := &proto.LoginRsp{}
	mapstructure.Decode(req.Body.Msg, loginReq)

	user := &model.User{}
	ok, err := db.Engine.Table(user).Where("username=?", loginReq.Username).Get(user)
	if err != nil {
		log.Println("user Query error:", err)
		return
	}

	if !ok {
		return
	}

	rsp.Body.Code = 0

	loginRsp.UId = 1
	loginRsp.Username = "admin"
	loginRsp.Session = "as"
	loginRsp.Password = ""
	rsp.Body.Msg = loginRsp
}
