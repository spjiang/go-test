package main

import (
	"fmt"
	"sync"
	"time"
)

type Issue struct {
	DeviceId   int
	DeviceName string
}

func main() {

	var wg sync.WaitGroup

	list := map[int][]Issue{}
	deviceList := []Issue{
		{1, "1_name_1"},
		{1, "1_name_2"},
		{1, "1_name_3"},
		{2, "1_name_1"},
		{2, "1_name_2"},
		{2, "1_name_3"},
		{3, "1_name_1"},
		{3, "1_name_2"},
		{3, "1_name_3"},
	}

	for _, v := range deviceList {
		list[v.DeviceId] = append(list[v.DeviceId], v)
	}

	for _, IssueList := range list {
		wg.Add(1)
		go handle(IssueList, &wg)
	}
	wg.Wait()

	for{

	}
}

func handle(deviceIssueList []Issue, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, v := range deviceIssueList {
		fmt.Println(v)
		time.Sleep(time.Duration(1) * time.Second)
	}
}


