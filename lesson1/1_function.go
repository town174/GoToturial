package main

//先定义名字，在定义类型；先定义参数，在定义返回值
func add(x, y int) int {
	return x + y
}

//支持返回多值
func swap(x, y string) (string, string) {
	return y, x
}

//支持命名返回值
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//变量列表， var 语句用于声明一个变量列表 ？？
var c, python, java int = 1, 2, 3

//匿名函数
var UnNameFunc = func(a, b int) int {
	return a + b
}
