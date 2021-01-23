package go_test

import (
	"errors"
	"fmt"
	"testing"
)

func TestErr(t *testing.T) {
	err := testErrHandle()
	if err != nil {
		fmt.Printf("test err %p v:%v\n", err, err)
	}
}

func testErrHandle() (err error) {

	fmt.Printf("test err-1 %p\n", err)

	s, err := returnErr()
	fmt.Printf("test err-2 %p\n", err)
	if err != nil {
		// return
	}
	fmt.Println(s)
	s1, err := returnErr()
	fmt.Printf("test err-3 %p\n", err)
	if err != nil {
		return
	}
	fmt.Println(s1)
	return
}
func returnErr() (string, error) {
	return "", errors.New("错误测试")
}
