package net

import (
	"log"
	"strings"
	"sync"
)

type HandlerFunc func(req *WsMsgReq, rsp *WsMsgRsp)

type MiddlewareFunc func(handlerFunc HandlerFunc) HandlerFunc

type group struct {
	mutex         sync.RWMutex
	prefix        string
	handlerMap    map[string]HandlerFunc
	middlewareMap map[string][]MiddlewareFunc
	middlewares   []MiddlewareFunc
}

func (g *group) AddRouter(name string, handlerFunc HandlerFunc, middlewares ...MiddlewareFunc) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.handlerMap[name] = handlerFunc
	g.middlewareMap[name] = middlewares
}

func (g *group) Use(middlewares ...MiddlewareFunc) {
	g.middlewares = append(g.middlewares, middlewares...)
}

func (g *group) exec(name string, req *WsMsgReq, rsp *WsMsgRsp) {
	h, ok := g.handlerMap[name]
	if !ok {
		/*
			ws://127.0.0.1:8004 为网关服务器，该服务的路由只有 *.*，所以 account.login 请求该服务器的时候会找不到login对应的handleFunc，因此 h == nil
			执行Group * 和 AddRouter * 后， prefix = *， name = * ，* 组 下有个 handlerMap[*]
			prefix.name -> *.*

		*/
		h, ok = g.handlerMap["*"]
		if !ok {
			log.Println("路由未定义")
		}
	}
	if ok {
		// 中间件  执行路由之前 需要执行中间件代码
		for i, l := 0, len(g.middlewares); i < l; i++ {
			h = g.middlewares[i](h)
		}
		mm, ok := g.middlewareMap[name]
		if ok {
			for i, l := 0, len(mm); i < l; i++ {
				h = mm[i](h)
			}
		}
		h(req, rsp)
	}
}

type Router struct {
	group []*group
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Group(prefix string) *group {
	g := &group{
		prefix:        prefix,
		handlerMap:    make(map[string]HandlerFunc),
		middlewareMap: make(map[string][]MiddlewareFunc),
	}
	r.group = append(r.group, g)
	return g
}

func (r *Router) Run(req *WsMsgReq, rsp *WsMsgRsp) {
	// req.Body.Name  路径 登录业务  account.login  (account组标识) login (路由标识)
	strs := strings.Split(req.Body.Name, ".")
	prefix := ""
	name := ""
	if len(strs) == 2 {
		prefix = strs[0]
		name = strs[1]
	}

	for _, g := range r.group {
		if g.prefix == prefix {
			g.exec(name, req, rsp)

		} else if g.prefix == "*" {
			g.exec(name, req, rsp)

		}
	}
}
