package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ziweiSgServer/constant"
	"ziweiSgServer/server/common"
	"ziweiSgServer/server/web/logic"
	"ziweiSgServer/server/web/model"
)

var DefaultAccountController = &AccountController{}

type AccountController struct {
}

func (a *AccountController) Register(c *gin.Context) {
	/*
		1. 获取参数
		2. 根据用户名 查询数据库 有：返回 没有：注册
		3. 告诉前端 注册成功
	*/
	rq := &model.RegisterReq{}
	err := c.ShouldBind(rq)
	if err != nil {
		log.Println("参数不合法", err)
		c.JSON(http.StatusOK, common.Error(constant.InvalidParam, "参数不合法"))
		return
	}

	err = logic.DefaultAccountLogic.Register(rq)
	if err != nil {
		log.Println("注册业务逻辑失败", err)
		c.JSON(http.StatusOK, common.Error(err.(*common.MyError).Code(), err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.Success(constant.OK, nil))
}
