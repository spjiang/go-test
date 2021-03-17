package main

import (
	"log"
	"time"
)

var ch chan struct{} = make(chan struct{}, 2)

func foo() {
	ch <- struct{}{}
	log.Println("foo() 000")
	ch <- struct{}{}
	log.Println("foo() 111")
	time.Sleep(5 * time.Second)
	log.Println("foo() 222")
	close(ch)
	log.Println("foo() 333")
}


func main() {
	var b struct{}
	log.Println("main() 111");
	go foo()
	log.Println("main() 222");
	a := <-ch
	log.Println("main() 333", a);
	b  = <-ch
	log.Println("main() 444", b);
	c := <-ch
	log.Println("main() 555", c);
}
