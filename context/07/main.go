package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go Test(ctx)
	cancel()
	fmt.Println("main")
	time.Sleep(5 * time.Second)

}

func Test(ctx context.Context) {
	fmt.Println("test start")
	select {
	case <-ctx.Done():
		fmt.Println("test")
		return
	}
}
