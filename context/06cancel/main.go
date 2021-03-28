package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("main 函数 开始...")
	go parent()
	time.Sleep(time.Second * 10)
	fmt.Println("main 函数 退出")
}
func parent() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	fmt.Println("父 协程 开始...")
	go child(ctx)
	time.Sleep(time.Second * 5)
	fmt.Println("5时间到了，关闭子协程")
	cancel()
	fmt.Println("演示结束")
}

func child(ctx context.Context) {
	var do bool
	for {
		select {
		case <-ctx.Done():
			fmt.Println("子 协程 接受停止信号...")
			return
		default:
			if do == false {
				fmt.Println("子协程运行中...")
				do = true
			}
		}
	}

}
