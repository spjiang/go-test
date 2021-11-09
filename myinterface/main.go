package main

import (
	"fmt"
)

type Walker interface {
	Walk()
}

type Swimer interface {
	Swim()
}

type dog struct {
}

func (d *dog) Walk() {
	fmt.Println("Dog avalible Walk")
}

type duck struct {
}

func (d *duck) Walk(){
	fmt.Println("Duck avalible Walk")
}

func (d *duck) Swim(){
	fmt.Println("Duck avalible Swim")
}

func Run() {
	var w interface{}
	w = new(duck)
	v,ok := w.(Swimer)
	fmt.Println(v,ok)
	v.Swim()

	v1,ok1 := w.(Walker)
	fmt.Println(v1,ok1)
	v1.Walk()

	v2,ok2 := w.(*duck)
	v2.Walk()
	v2.Swim()
	fmt.Println(v2,ok2)
	//注意：
	//* w变量虽然保存了duck结构体实力，但类型不是结构体，
	//所以也就不能调用duck接口体方法。
	//
	//* v变量是Swimer接口类型的变量，所以只能调用Swimer下的方法。
	//不能调用Walker接口下的方法，虽然duck结构体实现了Walker接口
	//
	//* 相应的v1是Walker接口类型的变量，只能调用Walker的方法，
	//不能调用Swimer的方法
	//
	//* 如果w接口类型的值要同时可以调用Walk和Swim,可以将w类型断言为结构体duck
	//的指针类型，如上面v2类型的断言
	//
	//* 上面的流程：
	//w接口类型赋值，
	//然后转w接口类型的值的类型为Swimer接口类型
	//然后转w接口类型的值的类型为Walker类型
	//然后转w接口类型的值的类型为结构体duck的指针类型

}
