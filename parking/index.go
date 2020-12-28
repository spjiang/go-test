package main

import (
	"fmt"
	"net/url"
)

func main() {
	str := "type=online&car_plate=京PF2K83&car_logo=未知&color=蓝色&car_color=未知&start_time=1608626255&park_id=1&camera_id=18030003f044&VehicleType=轿车&picture=&closeup_pic="
	arr, _ := url.ParseQuery(str)
	fmt.Println(arr)

}
