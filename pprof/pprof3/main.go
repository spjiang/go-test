package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

var container []byte
var container2 []byte

func func1_3() {
	fmt.Printf("handler1 %d\n", len(container))
}

func func1_2() {
	container = make([]byte, 200*1024*1024)
	for i := 0; i < 200*1024*1024; i++ {
		container[i] = 0
	}

	func1_3()
}

func func1_1() {
	func1_2()
}

func func2_3() {
	container2 = make([]byte, 100*1024*1024)
	for i := 0; i < 100*1024*1024; i++ {
		container2[i] = 0
	}
	fmt.Printf("handler2 %d\n", len(container2))
}

func func2_2() {
	func2_3()
}

func func2_1() {
	func2_2()
}

func handler1(w http.ResponseWriter, r *http.Request) {
	func1_1()
	runtime.GC()
}

func handler2(w http.ResponseWriter, r *http.Request) {
	func2_1()
	runtime.GC()
}

func main() {
	http.HandleFunc("/test1", handler1)
	http.HandleFunc("/test2", handler2)
	http.ListenAndServe(":6060", nil)
}