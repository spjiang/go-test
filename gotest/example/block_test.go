package example

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

/**

什么时候会死锁，如何避免

*/

// 第一种情形：无缓存能力的管道，自己写完自己读
// 不能单个协程自读写一个没有缓冲能力的管道。
// 解决办法很简单，开辟两条协程，一条协程写，一条协程读。
func TestBlock01(t *testing.T) {
	ch := make(chan int, 0)
	ch <- 666
	x := <-ch
	fmt.Println(x)
}

// 第二种情形：协程来晚了
// 我们可以看到，这条协程开辟在将数字写入到管道之后，因为没有人读，管道就不能写，然后写入管道的操作就一直阻塞。这时候你就有疑惑了，不是开辟了一条协程在读吗？但是那条协程开辟在写入管道之后，如果不能写入管道，就开辟不了协程。
func TestBlock02(t *testing.T) {
	ch := make(chan int, 0)
	ch <- 666
	x := <-ch
	fmt.Println(x)
}

// 第三种情形：管道读写时，相互要求对方先读/写
func TestBlock03(t *testing.T) {
	chHusband := make(chan int, 0)
	chWife := make(chan int, 0)

	go func() {
		select {
		case <-chHusband: // 先读chHusband在写chWife
			chWife <- 888
		}
	}()

	select {
	case <-chWife: // 先读chWife在写chHusband
		chHusband <- 888
	}
}

// 第四种情形：读写锁相互阻塞，形成隐形死锁
//这两条协程，如果第一条协程先抢到了只写锁，另一条协程就不能抢只读锁了，那么因为另外一条协程没有读，所以第一条协程就写不进。
//如果第二条协程先抢到了只读锁，另一条协程就不能抢只写锁了，那么因为另外一条协程没有写，所以第二条协程就读不到。
func TestBlock04(t *testing.T) {
	var rmw09 sync.RWMutex
	ch := make(chan int, 0)

	go func() {
		rmw09.Lock()
		ch <- 123
		rmw09.Unlock()
	}()

	go func() {
		rmw09.RLock()
		x := <-ch
		fmt.Println("读到", x)
		rmw09.RUnlock()
	}()

	for {
		runtime.GC()
	}
}
