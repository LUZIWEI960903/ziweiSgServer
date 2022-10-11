package controller

import (
	"strings"
	"sync"
	"ziweiSgServer/config"
	"ziweiSgServer/net"
)

var GateHandler = &Handler{
	proxyMap: make(map[string]map[int64]*net.ProxyClient),
}

type Handler struct {
	proxyMutex sync.Mutex
	proxyMap   map[string]map[int64]*net.ProxyClient // 代理地址 -> 客户端连接（游戏客户端id -> 连接）
	loginProxy string
	gameProxy  string
}

func (h *Handler) Router(r *net.Router) {
	h.loginProxy = config.File.MustValue("gate_server", "login_proxy", "ws://127.0.0.1:8003")
	h.gameProxy = config.File.MustValue("gate_server", "game_proxy", "ws://127.0.0.1:8001")
	g := r.Group("*")
	g.AddRouter("*", h.all)
}

func (h *Handler) all(req *net.WsMsgReq, rsp *net.WsMsgRsp) {
	//fmt.Println("网关的处理器")
	// account 转发
	name := req.Body.Name // name = "account.login"
	proxyStr := ""
	if isAccount(name) {
		proxyStr = h.loginProxy
	}
	proxy := net.NewProxyClient(proxyStr)
	proxy.Connect()
}

func isAccount(name string) bool {
	return strings.HasPrefix(name, "account.")
}
