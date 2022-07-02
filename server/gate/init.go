package gate

import (
	"ziweiSgServer/net"
	"ziweiSgServer/server/gate/controller"
)

var Router = net.NewRouter()

func Init() {
	initRouter()
}

func initRouter() {
	controller.GateHandler.Router(Router)
}
