package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/web"
	"jtthink/framework/gin_"
	_ "jtthink/src/lib"
	"log"
)

func main()  {

	r:=gin.New()

	gin_.BootStrap(r)

	//web改成micro 就是grpc,并直接注册到etcd里面
	service:=web.NewService(
		web.Name("api.jtthink.com.http.course"),
		web.Handler(r),
	 )

	service.Init()
	if err:= service.Run(); err != nil {
		log.Println(err)
	}

}