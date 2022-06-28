package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"ziweiSgServer/config"
	"ziweiSgServer/server/web"
)

func main() {
	host := config.File.MustValue("web_server", "host")
	port := config.File.MustValue("web_server", "port")

	router := gin.Default()
	// 路由
	web.Init(router)

	s := &http.Server{
		Addr:           fmt.Sprintf("%v:%v", host, port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	log.Println("web server error:", err)
}
