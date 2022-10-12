package controller

import "ziweiSgServer/net"

var DefaultRoleController = &RoleController{}

type RoleController struct {
}

func (r *RoleController) Router(router *net.Router) {
	g := router.Group("role")
	g.AddRouter("enterServer", r.enterServer)
}

func (r *RoleController) enterServer(req *net.WsMsgReq, rsp *net.WsMsgRsp) {
	// 进入游戏的逻辑
	// session 需要验证是否合法， 合法 取出登录用户的 id
	// 根据用户 id 查询对应的游戏角色，如果有 就继续，没有 提示无角色
	// 根据角色id 查询角色拥有的资源 roleRes，如果有则返回，没有 初始化资源

}
