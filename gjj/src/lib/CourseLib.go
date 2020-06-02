package lib

import (
	"github.com/gin-gonic/gin"
	"jtthink/framework/gin_"
	"jtthink/src/Course"
	"jtthink/src/service"
)

func init()  {
	courseService:=service.NewCourseServiceImpl()
	gin_.NewBuilder().WithService(courseService).
		 WithMiddleware(Check_Middleware()).
		WithEndpoint(Courselist_Endpoint(courseService)).
		WithRequest(Courselist_Request()).
		WithResponse(Course_Response()).Build("/test","GET")
}



//获取列表相关
func Courselist_Endpoint(c *service.CourseServiceImpl)   gin_.Endpoint {
	return func(context *gin.Context, request interface{}) (response interface{}, err error) {
		rsp:=&Course.ListResponse{}
	    err=c.ListForTop(context,request.(*Course.ListRequest),rsp)
	    return rsp,err
	}
}
//这个函数的作用是怎么处理请求
func Courselist_Request() gin_.EncodeRequestFunc{
	return func(context *gin.Context) (i interface{}, e error) {
		bReq:=&Course.ListRequest{}
		err:=context.BindQuery(bReq) //使用的是query 参数
		if err!=nil{
			return nil,err
		}
		return bReq,nil
	}
}
//这个函数作用是：怎么处理响应结果
func Course_Response()  gin_.DecodeResponseFunc  {
	return func(context *gin.Context, res interface{}) error {
		context.JSON(200,res)
		return nil
	}
}