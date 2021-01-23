package main

import (
	"net/http"
	"time"
)
import _ "net/http/pprof"

func main() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
	time.Sleep(100*time.Minute)
}