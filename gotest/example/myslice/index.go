package myslice

//定义一个函数，给切片添加一个元素
func AddOne(s []int) {
	s[0] = 4  // 可以改变原切片值
	s = append(s, 1)  // 扩容后分配了新的地址，原切片将不再受影响
	s[0] = 8
}
