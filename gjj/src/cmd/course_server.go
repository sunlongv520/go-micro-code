package main

import (
	"github.com/micro/go-micro/v2"
	"jtthink/src/Course"
	"jtthink/src/service"

	"log"
)


func main()  {
	cService:=micro.NewService(
		micro.Name("api.jtthink.com.course"))
	cService.Init()

	err:=Course.RegisterCourseServiceHandler(cService.Server(),service.NewCourseServiceImpl())
	if err!=nil{
		log.Fatal(err)
	}
	if err = cService.Run(); err != nil {
		log.Println(err)
	}
}