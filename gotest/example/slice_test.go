package example

import (
	"fmt"
	"gotest/example/myslice"
	"testing"
)

func TestSlice(*testing.T) {
	var s1 = []int{2}  // 初始化一个切片
	fmt.Println(s1)
	myslice.AddOne(s1) // 调用函数添加一个切片
	fmt.Println(s1)    // 输出一个值 [4]
}


func TestSlice01(*testing.T) {
	var slice []int
	slice = append(slice,1)
	fmt.Println(slice)
}