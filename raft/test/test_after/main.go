package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	election()
}

func randRange(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}

func election() {
	for {
		// 设置超时，150到300随机数
		timeout := randRange(150, 300)
		fmt.Printf("timeout:%d\n", timeout)
		select {
		// 延迟等待1毫秒
		case <-time.After(time.Duration(timeout) * time.Millisecond):
			fmt.Printf("time.After,timeout:%d\n", timeout)
		}
	}
}
