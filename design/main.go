package main

import (
	"fmt"
	"github.com/spjiang/go-test/design/facade"
)

func main() {
	a := facade.NewAPI()
	s := a.Test()
	fmt.Println(s)
}
