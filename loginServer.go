package main

import (
	"ziweiSgServer/config"
	"ziweiSgServer/net"
	"ziweiSgServer/server/login"
)

func main() {
	host := config.File.MustValue("login_server", "host", "127.0.0.1")
	port := config.File.MustValue("login_server", "port", "8003")

	s := net.NewServer(host + ":" + port)
	s.NeedSecret(false)
	login.Init()
	s.Router(login.Router)
	s.Start()

}
