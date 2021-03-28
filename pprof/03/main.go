package main

import (
	"fmt"
	"net/http"
	"time"
)
import _ "net/http/pprof"
// logicCode 一段有问题的代码
func logicCode() {
	var c chan int // nil
	for{
		select {
		case v:= <-c: // 阻塞
			fmt.Printf("recv from chan,value:%v\n",v)
		default:
			time.Sleep(500*time.Millisecond)
		}
	}
}

func main()  {
	for i := 0; i < 6; i++ {
		go logicCode()
	}
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
	for{}

}

