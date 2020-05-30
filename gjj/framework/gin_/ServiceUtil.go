package gin_

import "github.com/gin-gonic/gin"

var HandlerMap map[string]map[string]gin.HandlerFunc

func init()  {
	HandlerMap=make( map[string]map[string]gin.HandlerFunc)
}
func SetHandler(methods string,path string,handler gin.HandlerFunc)  {
	if _,ok:=HandlerMap[path];ok{
		HandlerMap[path][methods]=handler
	}else{
		HandlerMap[path]=make(map[string]gin.HandlerFunc)

	}
	HandlerMap[path][methods]=handler
}
func BootStrap(router *gin.Engine)  {
	for path,hmap:=range  HandlerMap{
		for method,handler:=range hmap{
			router.Handle(method,path,handler)
		}
	}
}