package main

import (
	"fmt"
	"time"
)
import _ "net/http/pprof"
import "net/http"

func worker(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("done")
}

//func main() {
//	ch := make(chan int)
//	go worker(ch)
//	// 发送完数据后忘了 close(ch)
//}

//	func main() {
//		ch := make(chan int)
//		go func() {
//			<-ch // 永远没人发送
//		}()
//		time.Sleep(time.Hour)
//	}
func main() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
	ch := make(chan int)
	go func() {
		ch <- 1 // 没人接收，永久阻塞
	}()
	time.Sleep(time.Hour)
}
