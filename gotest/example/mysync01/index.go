package mysync01

import (
	"fmt"
	"sync"
)

var (
	plateVal string     = ""           // 盘子
	plate    sync.Mutex = sync.Mutex{} // 盘子互斥锁
	apple    sync.Mutex = sync.Mutex{} // 苹果
	orange   sync.Mutex = sync.Mutex{} // 橘子
)

func init() {
	apple.Lock()
	orange.Lock()
}

func dad() { // 父亲
	for {
		plate.Lock()       // 获取盘子互斥锁
		plateVal = "apple" // 像盘子中放苹果
		apple.Unlock()     // 允许取苹果
	}
}

func mom() { // 母亲
	for {
		plate.Lock()        // 获取盘子互斥锁
		plateVal = "orange" // 向瓶子中放橘子
		orange.Unlock()     // 允许取橘子
	}
}

func son() { // 儿子
	for {
		orange.Lock()   // 互斥的取橘子
		get := plateVal // 从盘子中取橘子
		plateVal = ""
		plate.Unlock()              // 允许向盘子放东西
		fmt.Println("son eat", get) // 吃
	}
}

func daughter() {
	for {
		apple.Lock()    // 互斥的取苹果
		get := plateVal // 从盘子中取苹果
		plateVal = ""
		plate.Unlock()                   // 允许向盘子放东西
		fmt.Println("daughter eat", get) // 吃
	}
}
