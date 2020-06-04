package lib

import (
	"github.com/gin-gonic/gin"
	"jtthink/framework/gin_"
	"jtthink/src/Course"
	"jtthink/src/service"
)

func init()  {
	courseService:=service.NewCourseTopicServiceImpl()
	gin_.NewBuilder().WithService(courseService).
		WithMiddleware(Check_Middleware()).
		 WithEndpoint(CourseTopicList_Endpoint(courseService)).
		WithRequest(CourseTopicList_Request()).
		WithResponse(Course_Response()).Build("/topic/:cid","GET")


}

func CourseTopicList_Endpoint(c *service.CourseTopicServiceImpl)   gin_.Endpoint {
	return func(context *gin.Context, request interface{}) (response interface{}, err error) {
		rsp:=&Course.TopicResponse{Result:make([]*Course.CourseTopic,0)}
		err=c.GetTopic(context,request.(*Course.TopicRequest),rsp)
		return rsp,err
	}
}
//这个函数的作用是怎么处理请求
func CourseTopicList_Request() gin_.EncodeRequestFunc{
	return func(context *gin.Context) (i interface{}, e error) {
		bReq:=&Course.TopicRequest{}
		err:=context.BindUri(bReq) //使用的是query 参数
		if err!=nil{
			return nil,err
		}
		return bReq,nil
	}
}