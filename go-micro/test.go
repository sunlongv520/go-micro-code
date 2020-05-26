package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"log"
)



func callAPI2(s selector.Selector){
	var myClient = client.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
	)
	fmt.Println(myClient)
	req:=client.NewRequest("prodservice","/v1/prods",map[string]string{})
	var rsp map[string]interface{}
	err:=myClient.Call(context.Background(),req,&rsp)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(rsp["data"])
}


func main()  {
	etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))

	mySelector := selector.NewSelector(
		selector.Registry(etcdReg),
		selector.SetStrategy(selector.RoundRobin),
		)
	fmt.Println(mySelector)
	callAPI2(mySelector)


}