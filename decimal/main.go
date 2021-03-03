package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main()  {
	var f1 float64 = 0.3
	var f2 float64 = 0.6
	//fmt.Println(f1 + f2)

	//初始化一个小数
	price := decimal.NewFromFloat(f1)
	price2:= decimal.NewFromFloat(f2)
	total := price.Add(price2)
	fmt.Println(total)
}

