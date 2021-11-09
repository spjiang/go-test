package main

import "fmt"

func main()  {
	Test()
}

func Test() {

	user := map[int]string{
		1: "蒋",
		2: "李",
	}

	userData := map[string]map[int]string{
		"重庆": user,
	}

	fmt.Println(userData)
}

