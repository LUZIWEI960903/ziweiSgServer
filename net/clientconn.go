package net

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type ClientConn struct {
	wSConn        *websocket.Conn
	handshake     bool
	handshakeChan chan bool
}

func (c *ClientConn) Start() bool {
	//做的事情 就是 不停地接收消息
	// 等待握手的信息返回
	c.handshake = false
	go c.wsReadLoop()
	return c.waitHandshake()
}

func (c *ClientConn) waitHandshake() bool {
	// 等待握手成功  等待握手的消息
	// 万一程序出现问题了 一直响应不到  设置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	select {
	case _ = <-c.handshakeChan:
		log.Println("waitHandshake() success")
		return true
	case <-ctx.Done():
		log.Println("waitHandshake() timeout")
		return false
	}

}

func (c *ClientConn) wsReadLoop() {
	for {
		_, data, err := c.wSConn.ReadMessage()
		fmt.Println(data, err)
		//收到握手消息了
		c.handshake = true
		c.handshakeChan <- c.handshake
	}
}

func NewClientConn(wsConn *websocket.Conn) *ClientConn {
	return &ClientConn{
		wSConn:        wsConn,
		handshakeChan: make(chan bool),
	}
}
