package main

import (
	"fmt"
	"strconv"
)

func main()  {

	var ints = []int64{1,2,3,4}

	S := ""
	for _,i := range ints {
		S = S +","+ strconv.FormatInt(i,10)
	}

	fmt.Println(S)
}
