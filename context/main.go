package main

import (
	"errors"
	"fmt"
)

func main() {
	var err error
	fmt.Printf("err type %p\n", err)
	s, err := test()
	fmt.Printf("err type %p\n", err)
	fmt.Println(s)
	s2, err := test()
	fmt.Printf("err type %p\n", err)
	fmt.Println(s2)

}

func test() (string, error) {
	return "", errors.New("cc")
}
