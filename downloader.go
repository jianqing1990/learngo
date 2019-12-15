package main

import (
	"fmt"
	"learngo/infra"
)

func getRetriever() infra.Retriever {
	return infra.Retriever{}
}

func main() {
	var retriever infra.Retriever = getRetriever()
	//	retriever := getRetriever()
	fmt.Println(retriever.Get("https://www.imooc.com/"))
}
