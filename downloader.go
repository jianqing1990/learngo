package main

import (
	"fmt"
	//	"learngo/infra"
	"learngo/testing"
)

func getRetriever() testing.Retriever {
	return testing.Retriever{}
}

//?:Someing that can Get
type retriever interface {
	Get(string) string
}

func main() {
	//var retriever testing.Retriever = getRetriever()
	var r retriever = getRetriever()
	//	retriever := getRetriever()
	//fmt.Println(retriever.Get("https://www.imooc.com/"))
	fmt.Println(r.Get("https://www.imooc.com/"))

}
