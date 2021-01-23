package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

// logicCode 一段有问题的代码
func logicCode() {
	var c chan int // nil
	for{
		select {
		case v:= <-c: // 阻塞
			fmt.Printf("recv from chan,value:%v\n",v)
		default:
			time.Sleep(500*time.Millisecond)
		}
	}
}

func main() {
	var isCPUpprof bool // 是否开启CPUprofile的标志位
	var isMempprof bool // 是否开启Memprofile的标志位
	flag.BoolVar(&isCPUpprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMempprof, "mem", false, "turn cpu pprof on")
	flag.Parse()
	if isCPUpprof {
		fmt.Println("cpu--->\n")
		f1, err := os.Create("./cpu.profile")
		if err != nil {
			fmt.Printf("create cpu pprof failed,err:%v\n", err)
			return
		}
		_ = pprof.StartCPUProfile(f1)
		defer func() {
			pprof.StopCPUProfile()
			f1.Close()
		}()
	}
	for i := 0; i < 6; i++ {
		go logicCode()
	}
	time.Sleep(20 * time.Second)
	if isMempprof {
		f2, err := os.Create("./mem.profile")
		if err != nil {
			fmt.Printf("create mem pprof failed,err:%v\n", err)
			return
		}
		_ = pprof.WriteHeapProfile(f2)
		defer f2.Close()
	}

}
