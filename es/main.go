package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"log"
)

func main() {

	client,err:=elastic.NewClient(
		elastic.SetURL("http://192.168.29.135:9200/"),
		elastic.SetSniff(false),
		)
	if err!=nil{
		log.Fatal(err)
	}
	ctx:=context.Background()
	mapping,err:=client.GetMapping().Index("news").Do(ctx)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(mapping)
}
