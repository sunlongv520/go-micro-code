package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	_ "jtthink/src/Course"
	"jtthink/src/Users"
	"log"
)

type UserService struct{
	client client.Client
}
func (this *UserService) Test(ctx context.Context, req *Users.UserRequest, rsp *Users.UserResponse) error {
	rsp.Ret="users"+req.Id
	//c:=Course.NewCourseService("api.jtthink.com.course",this.client)
	//course_rsp,_:=c.ListForTop(ctx,&Course.ListRequest{Size:10})
	//log.Println(course_rsp.Result)
	log.Println(req.Id)
	log.Println(rsp.Ret)
	return nil
}
func NewUserService(c client.Client) *UserService  {
	return &UserService{client:c}
}
func main()  {

	service:=micro.NewService(
		micro.Name("api.jtthink.com.user"))
	service.Init()



	err:=Users.RegisterUserServiceHandler(service.Server(),NewUserService(service.Client()))
	if err!=nil{
		log.Fatal(err)
	}
	if err = service.Run(); err != nil {
		log.Println(err)
	}



}
