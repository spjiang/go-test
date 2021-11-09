package main

import (
	"fmt"
	"time"
)

func main(){

	ch := make(chan int, 4)

	ch <-1
	ch <-2
	ch <-3
	ch <-4

	close(ch)
	time.Sleep(5*time.Second)
	for value := range ch {
		fmt.Println("value:", value)
	}

}