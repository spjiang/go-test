package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//ctx, cancel := context.WithCancel(context.Background())
	//go test(ctx)
	//time.Sleep(5 * time.Second)
	//cancel()
	//time.Sleep(5 * time.Second)
	//fmt.Println("over")
	var count int = 20
	var countPoint *int
	countPoint = &count
	fmt.Printf("count 变量的地址：%x \n", &count)
	fmt.Printf("count 变量的地址：%p \n", &count)
	fmt.Printf("countPoint 变量存储的值内容：%v \n", *countPoint)
	fmt.Printf("countPoint 变量的地址：%x \n", &countPoint)
	fmt.Printf("countPoint 变量的地址：%p \n", &countPoint)
	fmt.Printf("count 变量的地址：%p \n", &count)
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
