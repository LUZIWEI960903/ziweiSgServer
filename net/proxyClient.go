package net

import (
	"errors"
	"github.com/gorilla/websocket"
	"time"
)

type ProxyClient struct {
	proxy string
	conn  *ClientConn
}

func (c *ProxyClient) Connect() error {
	//通过Dialer连接websocket服务器
	var dialer = websocket.Dialer{
		Subprotocols:     []string{"p1", "p2"},
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		HandshakeTimeout: 30 * time.Second,
	}
	ws, _, err := dialer.Dial(c.proxy, nil)
	if err == nil {
		c.conn = NewClientConn(ws)
		if !c.conn.Start() {
			return errors.New("握手失败")
		}
	}
	return err
}

func (c *ProxyClient) SetProperty(key string, data interface{}) {
	if c.conn != nil {
		c.conn.SetProperty(key, data)
	}
}

func (c *ProxyClient) SetOnPush(hookFunc func(conn *ClientConn, body *RspBody)) {
	if c.conn != nil {
		c.conn.SetOnPush(hookFunc)
	}
}

func (c *ProxyClient) Send(name string, msg interface{}) (*RspBody, error) {
	if c.conn != nil {
		return c.conn.Send(name, msg), nil
	}
	return nil, errors.New("连接未发现...")
}

func NewProxyClient(proxy string) *ProxyClient {
	return &ProxyClient{
		proxy: proxy,
	}
}
