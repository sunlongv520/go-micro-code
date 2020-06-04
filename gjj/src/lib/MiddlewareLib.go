package lib

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jtthink/framework/gin_"
	"net/http"
)

//放 通用中间件

//检查服务是否可用
func Check_Middleware()  gin_.Middleware {
	return func(next gin_.Endpoint) gin_.Endpoint {
		return func(context *gin.Context, request interface{}) (response interface{}, err error) {
			fmt.Println("我是中间件.....")
			return next(context,request)
		}
	}
}

func Cors_Middleware() gin_.Middleware {
	return func(next gin_.Endpoint) gin_.Endpoint {
		return func(c *gin.Context, request interface{}) (response interface{}, err error) {
			method := c.Request.Method
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")

			//放行所有OPTIONS方法
			if method == "OPTIONS" {
				c.AbortWithStatus(http.StatusNoContent)
			}
			// 处理请求
			return next(c,request)
		}
	}
}