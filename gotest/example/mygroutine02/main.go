package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var ObjectMap sync.Map

type Session struct {
	Id     int
	Achan  chan []byte
	Bchan  chan []byte
	ctx    context.Context
	cancel context.CancelFunc
}

func (s *Session) Start() {
	fmt.Printf("ID [%d] Start....\n", s.Id)
	go s.process1()
	go s.process2()
}

func (s *Session) process1() {
	defer func() {
		fmt.Printf("ID [%d] process1 defer....\n", s.Id)
	}()
	fmt.Printf("ID [%d] process1 start....\n", s.Id)
	for {
		select {
		case msg := <-s.Achan:
			fmt.Printf("ID [%d] process1 select....\n", s.Id)
			s.Bchan <- msg
		case <-s.ctx.Done():
			fmt.Printf("ID [%d] process1 ctx close....\n", s.Id)
			return
		}

	}
	fmt.Printf("ID [%d] Process1 end....\n", s.Id)
}

func (s *Session) process2() {

	defer func() {
		fmt.Printf("ID [%d] process2 defer....\n", s.Id)
	}()

	fmt.Printf("ID [%d] process2 start....\n", s.Id)
	time.Sleep(5 * time.Second)
	for msg := range s.Bchan {
		fmt.Printf("ID [%d] process2 range....\n", s.Id)
		go s.handle(string(msg))
	}
	fmt.Printf("ID [%d] Process end....\n", s.Id)
}

func (s *Session) handle(msg string) {
	fmt.Printf("ID [%d] handle msg[%s]....\n", s.Id, msg)
}

func (s *Session) Close() {
	fmt.Printf("ID [%d] Close start....\n", s.Id)
	close(s.Bchan)
	close(s.Achan)
	// s.cancel()
	ObjectMap.Delete(s.Id)
}

func A() {
	for i := 1; i <= 10; i++ {
		obj := new(Session)
		obj.Id = i
		obj.Achan = make(chan []byte, 100)
		obj.Bchan = make(chan []byte, 100)
		obj.ctx, obj.cancel = context.WithCancel(context.Background())
		ObjectMap.Store(i, obj)
		obj.Start()
		obj.Achan <- []byte(string(i))

	}

	time.Sleep(2 * time.Second)

	for i := 1; i <= 10; i++ {
		obj, ok := ObjectMap.Load(i)
		if ok {
			go obj.(*Session).Close()
		}
	}
}

func main() {

	go A()

	for {
		// 3. 遍历
		ObjectMap.Range(func(key, value interface{}) bool {
			k := key.(int)
			val := value.(*Session)
			fmt.Println(k, val)
			return true
		})
		fmt.Println("遍历...")
		time.Sleep(time.Second)
	}

}
