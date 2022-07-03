package main

import (
	"ziweiSgServer/config"
	"ziweiSgServer/net"
	"ziweiSgServer/server/gate"
)

/*
	1. 登录功能  account.login  需要通过网关 转发 登录服务器
	2. 网关 如何和 登录服务器 交互； 登录服务器作为websocket的服务端；网关作为websocket的客户端
	3. 网关又合游戏客户端交互；游戏客户端作为websocket的客户端；网关作为websocket的服务端
	4. websocket的服务端 已经实现了
	5. websocket的客户端 需要实现
	6. 网关：代理服务器（代理地址 代理的连接通道） 客户端连接（websocket连接）
	7. 路由：接收所有的请求；作为 网关的 websocket服务端的功能
	8. 握手协议 检测第一次建立连接的时候 授信
*/

func main() {
	host := config.File.MustValue("gate_server", "host", "127.0.0.1")
	port := config.File.MustValue("gate_server", "port", "8004")

	s := net.NewServer(host + ":" + port)
	gate.Init()
	s.Router(gate.Router)
	s.Start()
}
