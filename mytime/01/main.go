package main

import (
	"fmt"
	"time"
)

func main() {
	ctime := time.Now().Unix()
	nextDayTime := ctime + int64(86400)
	today := time.Unix(ctime, 0).Day()
	nextDay := time.Unix(nextDayTime, 0).Format("2016_")
	fmt.Println(today)
	fmt.Println(nextDay)

}
