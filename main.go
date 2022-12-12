package main

import "fmt"

func main() {
	a := 1
	p := &a //取址
	fmt.Println("%d\n", *p)
}
