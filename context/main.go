package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"
)

var c = 1

func doSome(i int) error {
	c++
	fmt.Println(c)
	if c > 3 {
		return errors.New("err occur")
	}
	return nil
}

func speakMemo(ctx context.Context, cancelFunc context.CancelFunc) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("ctx.Done")
			return
		default:
			fmt.Println("exec default func")
			err := doSome(3)
			if err != nil {
				fmt.Printf("cancelFunc()")
				cancelFunc()
			}
		}
	}
}

type A struct {
	AA AA
}

type AA struct {
	Name string
}

type B struct {
	Name string
}

func main() {
	//rootContext := context.Background()
	//ctx, cancelFunc := context.WithCancel(rootContext)
	//go speakMemo(ctx, cancelFunc)
	//time.Sleep(time.Second * 5)

	userdata := &A{
		AA: AA{
			Name: "name",
		},
	}

	B := &B{
		Name: userdata.AA.Name,
	}
	fmt.Println(B)

	// 2018-11-26 22:32:41
	t, _ := time.Parse("2006-01-02 15:04:05.000", "2018-11-26 14:32:41.998")
	alarmStartTime := t.Unix() - 15
	alarmEndTime := t.Unix() + 15

	ns := strconv.Itoa(t.Nanosecond())[0:3]

	startDateObj := time.Unix(alarmStartTime, 0)
	startDate := startDateObj.Format("2006-01-02 15:04:05") + "." + ns
	fmt.Println(startDate)

	endDateObj := time.Unix(alarmEndTime, 0)
	endDate := endDateObj.Format("2006-01-02 15:04:05") + "." + ns
	fmt.Println(endDate)

}
