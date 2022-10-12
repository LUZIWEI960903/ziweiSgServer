package main

import (
	"ziweiSgServer/config"
	"ziweiSgServer/net"
	"ziweiSgServer/server/game"
)

/*
	1. 登录完成后，创建角色
	2. 实际上就是根据用户 查询此用户拥有的角色；如果没有角色就 创建角色
	3. 木材 铁 令牌 金钱 主城 武将。。。 这些数据 要不要初始化， 已经玩过游戏，这些值是不是需要查询出来
	4. 地图 有关的，城池 资源土地 要塞。。。 需要定义
	5. 资源，军队，城池，武将。。。
*/

func main() {
	host := config.File.MustValue("game_server", "host", "127.0.0.1")
	port := config.File.MustValue("game_server", "port", "8001")

	s := net.NewServer(host + ":" + port)
	s.NeedSecret(false)
	game.Init()
	s.Router(game.Router)
	s.Start()
}
