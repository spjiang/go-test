package main

import (
	"fmt"
	"time"
)

func test()  {
	fmt.Println("ccc")
	fmt.Println(time.Now().Unix())
	time.Sleep(time.Duration(10)*time.Second)
}
func main()  {
	runTask(time.Duration(5)*time.Second,test)
	time.Sleep(time.Duration(10)*time.Minute)

}

func runTask(duration time.Duration, fn func()) {
	go func() {
		tricker := time.NewTicker(duration)
		defer tricker.Stop()
		for {
			fn()
			select {
			case <-tricker.C:
				fmt.Println(time.Now().Unix())
				fmt.Println("tricker")
			}
		}
	}()
}