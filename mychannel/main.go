package main

import (
	"fmt"
	"time"
)

type Task struct {
	DeviceId int
	DeviceName string
}




func Listen(ch <-chan Task) {
	//for task := range ch {
	//	fmt.Println(task)
	//}
	for  {
		select {
		case task:=<-ch:
			fmt.Println(task)
		}
	}
}

func addData(ch chan<- Task )  {
	for {
		time.Sleep(time.Duration(1)*time.Second)
		for i:=1;i<=3;i++ {
			t := time.Now().Unix()
			ch<-Task{
			i,
			"1_name_"+fmt.Sprintf("%d", t),
			}
		}
	}
}
func main() {

	ch := make(chan Task, 10)
	go addData(ch)
	go Listen(ch)
	deviceList := []Task{
		{1,"1_name_1"},
		{1,"1_name_2"},
		{1,"1_name_3"},
		{2,"1_name_1"},
		{2,"1_name_2"},
		{2,"1_name_3"},
		{3,"1_name_1"},
		{3,"1_name_2"},
		{3,"1_name_3"},
	}

	for _, v := range deviceList {
		device := v
		go func() {
			ch<-device
		}()
	}

	for{

	}
}
