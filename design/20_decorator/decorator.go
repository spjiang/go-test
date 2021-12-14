package decorator

import "fmt"

//装饰模式
//装饰模式使用对象组合的方式动态改变或增加对象行为。
//
//Go语言借助于匿名组合和非入侵式接口可以很方便实现装饰模式。
//
//使用匿名组合，在装饰器中不必显式定义转调原对象方法。

type Component interface {
	Calc() int
}

type ConcreteComponent struct{}

func (*ConcreteComponent) Calc() int {
	fmt.Println("ConcreteComponent-Calc")
	return 0
}

func WarpMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		Component: c,
		num:       num,
	}
}

type MulDecorator struct {
	Component
	num int
}

func (d *MulDecorator) Calc() int {
	fmt.Println("MulDecorator-Calc")
	return d.Component.Calc() * d.num
}

func WarpAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		Component: c,
		num:       num,
	}
}

type AddDecorator struct {
	Component
	num int
}

func (d *AddDecorator) Calc() int {
	fmt.Println("AddDecorator-Calc")
	return d.Component.Calc() + d.num
}
