package main

import "fmt"

func main() {
	var i = 10
	fmt.Println("i的内存地址:", &i)
	var ptr *int = &i
	fmt.Printf("ptr=%v \n", ptr)
	fmt.Printf("ptr本身内存地址%v \n", &ptr)
	fmt.Printf("ptr本身内存地址%p \n", ptr)
	

}
