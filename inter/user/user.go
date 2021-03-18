package user

import "fmt"

type User interface {
	UserInfo()
}

func Test() {
	fmt.Println("user test")
}
