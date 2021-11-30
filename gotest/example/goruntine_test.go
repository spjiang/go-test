package example

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine01(t *testing.T) {
	jobs := make(chan int, 100)
	results := make(chan int, 1)
	// 开启worker
	for w := 1; w <= 3; w++ {
		go Goroutine01Worker(w, jobs, results)
	}
	// 5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		fmt.Println(<-results)
	}
}

func Goroutine01Worker(workerId int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job: %d\n", workerId, j)
		time.Sleep(time.Second)
		//fmt.Printf("worker:%d end job: %d\n", workerId, j)
		results <- j * 2
	}
}
