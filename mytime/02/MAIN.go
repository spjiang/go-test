package main

import (
	"fmt"
	"time"
)

func main() {
	ctime := time.Now().Unix()

	startTime := ctime - 25920000
	data := GetMonthList(startTime,"2006_01")
	for _, month := range  data {
		fmt.Println(month)
	}
}


func GetMonthList(startTimestamp int64, format string) (data []string){
	defer func() {
		data = RemoveRepeatedElement(data)
	}()
	tmpTimestamp := startTimestamp
	data = append(data, "2020_12")
	for {
		if tmpTimestamp <= 0 {
			return
		}
		if tmpTimestamp > time.Now().Unix() {
			return
		}
		data = append(data, time.Unix(tmpTimestamp, 0).Format(format))
		tmpTimestamp += 28*24*60*60
	}
}

func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}