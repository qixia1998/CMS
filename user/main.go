package main

import (
	srv "common"
	"github.com/gin-gonic/gin"
	"user/router"
)

func main() {
	r := gin.Default()
	//路由
	router.InitRouter(r)
	//grpc服务注册
	gc := router.RegisterGrpc()
	//grpc服务注册到etcd
	router.RegisterEtcdServer()

	stop := func() {
		gc.Stop()
	}
	srv.Run(r, config.C.SC.Name, config.C.SC.Addr, stop)
}
