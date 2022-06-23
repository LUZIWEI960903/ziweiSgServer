package net

import (
	"encoding/json"
	"errors"
	"github.com/forgoer/openssl"
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"ziweiSgServer/utils"
)

// websocket服务
type wsServer struct {
	wsConn       *websocket.Conn
	router       *Router
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

func (w *wsServer) Router(router *Router) {
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
		// 收到消息 解析消息 前端发送过来的消息 就是json格式
		// 1. data 解压 unzip
		data, err = utils.UnZip(data)
		if err != nil {
			log.Println("UnZip(data) error", err)
			continue
		}
		// 2. 前端的消息的加密消息 进行解密
		secretKey, err := w.GetProperty("secretKey")
		if err == nil {
			//有加密
			key := secretKey.(string)
			//客户端传过来的数据是加密的 需要解密
			d, err := utils.AesCBCDecrypt(data, []byte(key), []byte(key), openssl.ZEROS_PADDING)
			if err != nil {
				log.Println("utils.AesCBCDecrypt error:", err)
				//出错后 发起握手
				w.Handshake()
			} else {
				data = d
			}
		}
		// 3. data 转为 body
		body := &ReqBody{}
		err = json.Unmarshal(data, body)
		if err != nil {
			log.Println("json.Unmarshal(data, body) error:", err)
		} else {
			// 获取到前端传递的数据了，拿上这些数据，去具体的业务进行处理
			req := &WsMsgReq{
				Body: body,
				Conn: w,
			}
			rsp := &WsMsgRsp{
				Body: &RspBody{
					Seq:  req.Body.Seq,
					Name: body.Name,
				},
			}
			w.router.Run(req, rsp)

			w.outChan <- rsp
		}
	}
	w.Close()
}

func (w *wsServer) writeMsgLoop() {
	for {
		select {
		case msg := <-w.outChan:
			w.Write(msg)
		}
	}
}

func (w *wsServer) Write(msg *WsMsgRsp) {
	// 发给客户端的数据转json
	data, err := json.Marshal(msg.Body)
	if err != nil {
		log.Println("json.Marshal(msg) error:", err)
	}
	secretKey, err := w.GetProperty("secretKey")
	if err == nil {
		//有加密
		key := secretKey.(string)
		//数据做加密
		data, _ = utils.AesCBCEncrypt(data, []byte(key), []byte(key), openssl.ZEROS_PADDING)
		//压缩
		if data, err := utils.Zip(data); err == nil {
			w.wsConn.WriteMessage(websocket.BinaryMessage, data)
		}
	}
}

func (w *wsServer) Close() {
	_ = w.wsConn.Close()
}

// 当游戏客户端 发送请求时 先会进行握手协议
// 后端会发送对应的加密key给客户端
// 客户端在发送数据时 会使用此key进行加密处理
const HandshakeName = "handshake"

func (w *wsServer) Handshake() {
	secretKey := ""
	key, err := w.GetProperty("secretKey")
	if err == nil {
		secretKey = key.(string)
	} else {
		secretKey = utils.RandSeq(16)
	}

	handshake := &Handshake{Key: secretKey}

	body := &RspBody{Name: HandshakeName, Msg: handshake}
	if data, err := json.Marshal(body); err == nil {
		if secretKey != "" {
			w.SetProperty("secretKey", secretKey)
		} else {
			w.RemoveProperty("secretKey")
		}
		if data, err = utils.Zip(data); err == nil {
			w.wsConn.WriteMessage(websocket.BinaryMessage, data)
		}
	}

}
