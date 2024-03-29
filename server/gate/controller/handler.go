package controller

import (
	"log"
	"strings"
	"sync"
	"ziweiSgServer/config"
	"ziweiSgServer/constant"
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
	} else {
		proxyStr = h.gameProxy
	}
	if proxyStr == "" {
		rsp.Body.Code = constant.ProxyNotInConnect
		return
	}
	//先去proxyMap中查询 是否有连接
	h.proxyMutex.Lock()
	_, ok := h.proxyMap[proxyStr]
	if !ok {
		h.proxyMap[proxyStr] = make(map[int64]*net.ProxyClient)
	}
	h.proxyMutex.Unlock()

	// 获取客户端id
	c, err := req.Conn.GetProperty("cid")
	if err != nil {
		log.Println("cid为获取到", err)
		rsp.Body.Code = constant.InvalidParam
		return
	}
	cid := c.(int64)
	proxy, ok := h.proxyMap[proxyStr][cid]
	if !ok {
		//没有建立连接 并发起调用
		proxy = net.NewProxyClient(proxyStr)
		h.proxyMutex.Lock()
		h.proxyMap[proxyStr][cid] = proxy
		h.proxyMutex.Unlock()

		err := proxy.Connect()
		if err != nil {
			h.proxyMutex.Lock()
			delete(h.proxyMap[proxyStr], cid)
			h.proxyMutex.Unlock()
			rsp.Body.Code = constant.ProxyConnectError
			return
		}
		proxy.SetProperty("cid", cid)
		proxy.SetProperty("proxy", proxyStr)
		proxy.SetProperty("gateConn", req.Conn)
		proxy.SetOnPush(h.onPush)
	}
	rsp.Body.Seq = req.Body.Seq
	rsp.Body.Name = req.Body.Name
	r, err := proxy.Send(req.Body.Name, req.Body.Msg)
	if err == nil {
		rsp.Body.Code = r.Code
		rsp.Body.Msg = r.Msg
	} else {
		rsp.Body.Code = constant.ProxyConnectError
		return
	}
}

func (h *Handler) onPush(conn *net.ClientConn, body *net.RspBody) {
	gc, err := conn.GetProperty("gateConn")
	if err != nil {
		log.Println("onPush gateConn,", err)
		return
	}
	gateConn := gc.(net.WSConn)
	gateConn.Push(body.Name, body.Msg)
}

func isAccount(name string) bool {
	return strings.HasPrefix(name, "account.")
}
