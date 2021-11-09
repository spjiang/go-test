package main

import (
	"fmt"
	"time"
)

func main()  {
	t := time.Now().Unix()
	fmt.Println(t)

	// 获取月份
	year := time.Now().Format("2006_01")
	fmt.Println(year)

	// 获取月份
	month := time.Now().Format("01")
	fmt.Println(month)

	// 获取月份
	day := time.Now().Day()
	fmt.Println(day)

	s, err := FormatDateToString("2021-07-06 10:00:02", "2006-01-02 15:04:05", "2006-01-02T15:04:05")
	fmt.Println(s,err)
}

func FormatDateToString(datePut string, putTimeTemplate string,outTimeTemplate string) (string, error) {
	stamp, err := time.ParseInLocation(putTimeTemplate, datePut, time.Local)
	if err != nil {
		return "", err
	}
	t := time.Unix(stamp.Unix(), 0)
	return t.Format(outTimeTemplate), nil
}