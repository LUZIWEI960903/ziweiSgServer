package net

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

// websocket服务
type wsServer struct {
	wsConn       *websocket.Conn
	router       *router
	outChan      chan *WsMsgRsp // 写队列
	Seq          int64
	property     map[string]interface{}
	propertyLock sync.RWMutex
}

func NewWsServer(wsConn *websocket.Conn) *wsServer {
	return &wsServer{
		wsConn:   wsConn,
		outChan:  make(chan *WsMsgRsp, 1000),
		Seq:      0,
		property: make(map[string]interface{}),
	}
}

func (w *wsServer) Router(router *router) {
	w.router = router
}

func (w *wsServer) SetProperty(key string, value interface{}) {
	w.propertyLock.Lock()
	defer w.propertyLock.Unlock()
	w.property[key] = value
}

func (w *wsServer) GetProperty(key string) (interface{}, error) {
	w.propertyLock.RLock()
	defer w.propertyLock.RUnlock()
	value, ok := w.property[key]
	if !ok {
		return nil, errors.New("GetProperty error")
	}
	return value, nil
}

func (w *wsServer) RemoveProperty(key string) {
	w.propertyLock.Lock()
	defer w.propertyLock.Unlock()
	delete(w.property, key)
}

func (w *wsServer) Addr() string {
	return w.wsConn.RemoteAddr().String()
}

func (w *wsServer) Push(name string, data interface{}) {
	w.outChan <- &WsMsgRsp{
		Body: &RspBody{
			Seq:  0,
			Name: name,
			Msg:  data,
		},
	}
}

// 通道一旦建立，那么收发消息 就得要一直监听才行
func (w *wsServer) Start() {
	// 启动读写数据的处理逻辑
	go w.readMsgLoop()
	go w.writeMsgLoop()
}

func (w *wsServer) readMsgLoop() {
	// 先读到客户端发送过来的消息，然后进行处理，然后再回消息
	// 经过路由，实际处理程序
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			w.Close()
		}
	}()

	for {
		_, data, err := w.wsConn.ReadMessage()
		if err != nil {
			log.Println("readMsgLoop() error", err)
			break
		}
		fmt.Printf("data: %v\n", data)
	}
	w.Close()
}

func (w *wsServer) writeMsgLoop() {
	for {
		select {
		case msg := <-w.outChan:
			fmt.Printf("msg: %v\n", msg)
		}
	}
}

func (w *wsServer) Close() {
	_ = w.wsConn.Close()
}
