package example

import (
	"fmt"
	"runtime"
	"testing"
)

func TestCpu01(t *testing.T)  {
	n := runtime.GOMAXPROCS(1) //将cpu设置为1核
	fmt.Println(n)
	n = runtime.GOMAXPROCS(1) //将cpu设置为1核
	fmt.Println(n)
}

func TestCpu02(t *testing.T)  {
	cpu02()
}
