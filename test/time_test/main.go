package main

import (
	"fmt"
	"github.com/spjiang/go-test/test/time_test/parse_time"
)

func main() {
	dateTime := "20080102150405"
	t, _ := parse_time.StringToTime("20060102150405", dateTime)
	r := t.Format("2006-01-02 15:04:05.000")
	fmt.Println(r)
}
