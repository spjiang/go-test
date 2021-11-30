package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 10)
	for i := 1; i < 100; i++ {
		fmt.Println("+++++++worker+++Id", i)
		go worker(jobs,i)
	}

	for i := 1; i < 1000000000; i++ {
		jobs <- i
		time.Sleep(time.Second * 1)
	}

	for {

	}
}

func worker(jobs <-chan int, workerId int) {
	for val := range jobs {
		time.Sleep(time.Second * 2)
		fmt.Printf("写入数据库%d....workerId:%d \r\n", val, workerId)
	}
}
