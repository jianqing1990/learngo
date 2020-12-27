package main

import (
	"learngo/test"
	"learngo/testing"
	"log"
)

func getRetriever() retriever {
	return testing.Retriever{}
}

//?:Someing that can Get
type retriever interface {
	Get(string) string
}

func main() {
	//var retriever testing.Retriever = getRetriever()
	//var r retriever = getRetriever()
	//	retriever := getRetriever()
	//fmt.Println(retriever.Get("https://www.imooc.com/"))
	//fmt.Println(r.Get("https://www.imooc.com/"))

	result := "sda"

	for {

		go func(i interface{}) {

			test.Item() <- result
			var sout func() chan interface{} = test.Item
			item := <-sout
			log.Printf("获取数据："+"%v", item)
		}(result)

	}

}
