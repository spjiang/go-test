package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go test(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(5 * time.Second)
	fmt.Println("over")
}

func test(ctx context.Context) {
	fmt.Println("test start...")
	select {
	case <-ctx.Done():
		fmt.Println("test exit...")
		return
	default:
	}
	fmt.Println("test running...")
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("+++++++++")
	}
}
