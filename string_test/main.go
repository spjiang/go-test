package main

import (
	"fmt"
	"strings"
)

func main() {
	//str := "1dc8969e23947b1b5ea504397d4cb80565e246330bd4813135e12e21229552701619329771"
	//l := len(str)
	//content := str[l-10 : ]
	//fmt.Println(content)

	//str := "camera/26.flv?sign=a17d86bf223dfdd0c2fb7548be173cb9459ed3c00fef710b58e703b7061fd6d61619330089"
	//sep:="?"
	//arr:=strings.Split(str,sep)
	//fmt.Println("arr:",arr[0])
	//fmt.Println(len(arr))


	str2 := "camera/26.flv?sign=a17d86bf223dfdd0c2fb7548be173cb9459ed3c00fef710b58e703b7061fd6d61619330089"
	stringPath := strings.TrimLeft(str2, "/")
	fmt.Println(stringPath)

	if strings.HasSuffix(stringPath, ".flv") {
		stringPath = strings.TrimRight(stringPath, ".flv")
	}

	fmt.Println(stringPath)


}