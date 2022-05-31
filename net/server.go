package net

import "net/http"

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

func (s *server) wsHandler(w http.ResponseWriter, r *http.Request) {

}
