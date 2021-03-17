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
	// close(notify)
	time.Sleep(time.Second * 5)
}
func test(s string, notify c) {
	fmt.Printf("start %s\n", s)
	for {
		select {
		case <-notify:
			fmt.Printf("end %s\n", s)
			return
		case <-time.After(1 * time.Second):
		}
	}
}
