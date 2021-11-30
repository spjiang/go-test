package mysync

import (
	"fmt"
	"sync"
	"time"
)

var n int = 10 // 缓冲区大小
var mutex sync.Mutex = sync.Mutex{} // 全局变量
var buf []int // 缓冲区

func init() {
	buf = make([]int,n)
	buf = buf[:0]
}

// 生产者
func productor(data int) {
	for {
		mutex.Lock()         // 获得锁 进入临界区
		for len(buf) == 10 { // 判断条件 缓冲区是否满了
			mutex.Unlock()          // 满的缓冲区 要将锁释放出来
			time.Sleep(time.Second) // 等一会儿
			mutex.Lock()            // 重新获取锁
		}

		buf = append(buf, data) // 将生产的数据写入缓冲区

		mutex.Unlock() // 释放锁
	}
}

// 消费者
func consumer() {
	for {
		mutex.Lock()        // 获得锁 进入临界区
		for len(buf) == 0 { // 判断条件 缓冲区是否为空
			mutex.Unlock()          // 空的缓冲区 要将锁释放出来
			time.Sleep(time.Second) // 等一会儿
			mutex.Lock()            // 重新获取锁
		}

		data := buf[len(buf)-1] // 从缓冲区获取数据
		buf = buf[:len(buf)-1]
		fmt.Println(data) // 打印到标准输出

		mutex.Unlock() // 释放锁
	}
}