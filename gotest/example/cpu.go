package example

import (
	"fmt"
	"runtime"
)

func cpu01() {
	//返回值n代表我们上次设置的cpu的最大个数
	n := runtime.GOMAXPROCS(1) //将cpu设置为1核
	fmt.Println(n)
	//将cpu设置为1核时，下面两个进程将会互相争夺，谁抢到，谁就会开始不断打印
	//所以打印结果是大片的1或0(谁抢到打印谁)，下面紧接着打印另外一个数
	//相比较给予比进程数多的核来讲，打印的数更加密集，因为多个核，两个进程竞争变小
	for {
		go fmt.Print(0) //子go程
		fmt.Print(1)    //主go程
	}
}

func cpu02() {
	//返回值n代表我们上次设置的cpu的最大个数
	n := runtime.GOMAXPROCS(runtime.NumCPU()) //将cpu设置为1核
	fmt.Println(n)
	//将cpu设置为1核时，下面两个进程将会互相争夺，谁抢到，谁就会开始不断打印
	//所以打印结果是大片的1或0(谁抢到打印谁)，下面紧接着打印另外一个数
	//相比较给予比进程数多的核来讲，打印的数更加密集，因为多个核，两个进程竞争变小
	for {
		go fmt.Print(0) //子go程
		fmt.Print(1)    //主go程
	}
}