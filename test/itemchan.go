package test

import "log"

func Item() chan interface{} {

	out := make(chan interface{})

	go func() {
		for {
			item := <-out
			log.Printf("获取数据："+"%v", item)
		}
	}()
	return out
}
