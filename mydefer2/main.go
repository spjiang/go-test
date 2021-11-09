package main

import "fmt"

func main()  {
	Test()
}

func Test()  {
	var err error
	defer func() {
		fmt.Println(err)
		fmt.Println()
	}()

	err = fmt.Errorf("Error 111")
}