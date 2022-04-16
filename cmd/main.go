package main

import (
	"mozzarella-login-center/dao"
	"mozzarella-login-center/route"
	"mozzarella-login-center/rpc"
)

func main() {
	dao.InitMysql()
	dao.InitRedis()
	go rpc.InitRpc()
	rpc.RegisterCenter()
	route.InitRoute()
}
