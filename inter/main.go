package main

import (
	"fmt"
	u "github.com/spjiang/go-test/inter/user"
)

type UserImpl struct {
}

func (u *UserImpl) UserInfo() {
	fmt.Println("test")
}

func Test2(user u.User) {
	user.UserInfo()
}

var ui = UserImpl{}

func main() {
	u.Test()

	Test2(&ui)

}
