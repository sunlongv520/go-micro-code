package main

import (
	"context"
	"github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-micro/v2/web"
	"jtthink/src/Course"
	"log"
	"net/http"
)

func main()  {
	service:=web.NewService(
		web.Name("api.jtthink.com.http.course"))
	c:=grpc.NewClient()

	service.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		c:=Course.NewCourseService("api.jtthink.com.course",c)
		course_rsp,_:=c.ListForTop(context.Background(),&Course.ListRequest{Size:10})
		log.Println(course_rsp.Result)
		writer.Write([]byte("http api test"))
	})
	service.Init()
	if err:= service.Run(); err != nil {
		log.Println(err)
	}
}