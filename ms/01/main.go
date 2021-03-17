package main

import (
	"fmt"
	"sync"
)

const N int = 10

func main() {
	m := make(map[int]int)
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			fmt.Println("+++",i)
			fmt.Println("+++2",m)
			fmt.Println()
			m[i] = i
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Println(m)
}
