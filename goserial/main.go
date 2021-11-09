//project main.go
package main

import (
	"fmt"
	"strings"

	"github.com/tarm/goserial"
)

const MAXRWLEN = 8000

func main() {

	cfg := &serial.Config{
		Name: "COM8",
		Baud: 115200,
		ReadTimeout: 3 /*毫秒*/}

	iorwc, err := serial.OpenPort(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer iorwc.Close()
	buffer := make([]byte, MAXRWLEN)

	//发命令之前清空缓冲区
	num, err := iorwc.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	//发命令数据类型为[]byte
	num, err = iorwc.Write([]byte("AT\r\n"))
	if err != nil {
		fmt.Println(err)
		return
	}

	var tmpstr string = ""
	for i := 0; i < 3000; i++ {

		num, err = iorwc.Read(buffer)
		if num > 0 {
			tmpstr += fmt.Sprintf("%s", string(buffer[:num]))
		}

		//查找读到信息的结尾标志
		if strings.LastIndex(tmpstr, "\r\nOK\r\n") > 0 {
			break
		}
	}

	//打印输出读到的信息
	fmt.Println(tmpstr)

	return
}