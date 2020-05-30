package gin_

type ServiceBuilder struct {
	service interface{}
	endPoint Endpoint
	requestFunc EncodeRequestFunc
	responseFunc DecodeResponseFunc
	methods string
	middlewares []Middleware

}

func NewBuilder() *ServiceBuilder  {
	 return &ServiceBuilder{middlewares:make([]Middleware,0)}

}

func(this *ServiceBuilder) WithService(obj interface{}) *ServiceBuilder {
	this.service=obj
	return this
}
func(this *ServiceBuilder) WithMiddleware(obj Middleware) *ServiceBuilder {
	this.middlewares=append(this.middlewares,obj)
	return this
}
func(this *ServiceBuilder) WithEndpoint(obj Endpoint) *ServiceBuilder {
	this.endPoint=obj
	return this
}
func(this *ServiceBuilder) WithRequest(obj EncodeRequestFunc) *ServiceBuilder {
	this.requestFunc=obj
	return this
}
func(this *ServiceBuilder) WithResponse(obj DecodeResponseFunc) *ServiceBuilder {
	this.responseFunc=obj
	return this
}
func(this *ServiceBuilder) WithMethods(methods string) *ServiceBuilder {
	this.methods=methods
	return this
}
func(this *ServiceBuilder) Build(path string,methods string)  *ServiceBuilder{

	finalEndpoint:=this.endPoint
	for _,m:=range this.middlewares{ //遍历中间件
		finalEndpoint=m(finalEndpoint)
	}
	 handler:=RegisterHandler(finalEndpoint,this.requestFunc,this.responseFunc)

	SetHandler(methods,path,handler)
	return this
}
