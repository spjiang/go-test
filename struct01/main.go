package main

import (
	"fmt"
	"time"
)

type c chan struct{}

func main() {
	notify := make(c)
	go test("01", notify)
	go test("02", notify)
	close(notify)
	time.Sleep(time.Second * 5)
}
func test(s string, notify c) {
	fmt.Printf("start %s\n", s)
	do := false
	for {
		select {
		case <-notify:
			fmt.Printf("end %s\n", s)
			return
		case <-time.After(1 * time.Second):
			//fmt.Println("+++")
			if do == false {
				fmt.Println("业务处理中...")
			}
			do = true
		}
	}
}
