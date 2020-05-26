package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/web"
	"gomicro.jtthink.com/ProdService"
)

func main(){
	 etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))

	ginRouter:=gin.Default()
	v1Group:=ginRouter.Group("/v1")
	{
		v1Group.Handle("GET","/prods", func(context *gin.Context) {
			context.JSON(200,
				gin.H{
					"data":ProdService.NewProdList(2)})
		})
	}
	server:=web.NewService(
		web.Name("prodservice"),
		//web.Address(":8001"),
		web.Handler(ginRouter),
		web.Registry(etcdReg),
		)

	server.Init()
	server.Run()

}