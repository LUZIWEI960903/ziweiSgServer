package controller

import (
	"github.com/mitchellh/mapstructure"
	"log"
	"ziweiSgServer/constant"
	"ziweiSgServer/db"
	"ziweiSgServer/net"
	"ziweiSgServer/server/login/model"
	"ziweiSgServer/server/login/proto"
	"ziweiSgServer/utils"
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
		// 没查出来
		rsp.Body.Code = constant.UserNotExist
		return
	}

	pwd := utils.Password(loginReq.Password, user.Passcode)
	if user.Passwd != pwd {
		rsp.Body.Code = constant.PwdIncorrect
		return
	}

	tokenStr, _ := utils.Award(user.UId)

	rsp.Body.Code = constant.OK

	loginRsp.UId = user.UId
	loginRsp.Username = user.Username
	loginRsp.Session = tokenStr
	loginRsp.Password = ""
	rsp.Body.Msg = loginRsp
}
