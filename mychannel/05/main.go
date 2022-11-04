package main

import (
	"fmt"
	"os"
)

func main() {
	fname := "./t.txt"
	f, err := os.OpenFile(fname, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	f.Write([]byte("test"))
	f.Close()
}
