package logic

import (
	"log"
	"time"
	"ziweiSgServer/constant"
	"ziweiSgServer/db"
	"ziweiSgServer/server/common"
	"ziweiSgServer/server/models"
	"ziweiSgServer/server/web/model"
	"ziweiSgServer/utils"
)

var DefaultAccountLogic = &AccountLogic{}

type AccountLogic struct {
}

func (*AccountLogic) Register(rq *model.RegisterReq) error {
	username := rq.Username
	user := &models.User{}
	ok, err := db.Engine.Table(user).Where("username=?", username).Get(user)
	if err != nil {
		log.Println("注册查询失败", err)
		return common.NewError(constant.DBError, "数据库异常")
	}

	if ok {
		return common.NewError(constant.UserExist, "用户已存在")
	} else {
		// 注册
		user.Username = rq.Username
		user.Passcode = utils.RandSeq(6)
		user.Passwd = utils.Password(rq.Password, user.Passcode)
		user.Hardware = rq.Hardware
		user.Ctime = time.Now()
		user.Mtime = time.Now()
		_, err := db.Engine.Table(user).Insert(user)
		if err != nil {
			log.Println("注册插入失败", err)
			return common.NewError(constant.DBError, "数据库异常")
		}
		return nil
	}
}
