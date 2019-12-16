package main

import (
	"fmt"
	//	"learngo/infra"
	"learngo/testing"
)

func getRetriever() testing.Retriever {
	return testing.Retriever{}
}

func main() {
	var retriever testing.Retriever = getRetriever()

	//	retriever := getRetriever()
	fmt.Println(retriever.Get("https://www.imooc.com/"))
}
