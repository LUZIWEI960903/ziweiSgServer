package net

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type server struct {
	addr   string
	router *router
}

func NewServer(addr string) *server {
	return &server{
		addr: addr,
	}
}

func (s *server) Start() {
	http.HandleFunc("/", s.wsHandler)
	err := http.ListenAndServe(s.addr, nil)
	if err != nil {
		panic(err)
	}
}

// http升级websocket协议的配置
var wsUpgrader = websocket.Upgrader{
	// 运行所有CORS跨域请求
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (s *server) wsHandler(w http.ResponseWriter, r *http.Request) {
	// http协议升级websocket协议
	wsConn, err := wsUpgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal("websocket connecting failed...", err)
	}
	log.Fatal("websocket connecting success...")
}
