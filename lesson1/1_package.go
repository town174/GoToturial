package main

import "fmt"

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

//11 基本类型

func main() {
	//fmt.Println("My Favorite number is", rand.Intn(20))

	//fmt.Println(rand.pi)

	//fmt.Println(add(5, 8))
	//a, b := swap("5", "8")
	//fmt.Println(a, b)

	//fmt.Println(split(87))

	//var i int
	//fmt.Println(i, c, python, java)

	//短变量声明:= 可在类型明确的地方代替 var 声明
	//不能函数外使用
	k, m := 3, 4
	fmt.Println(k, m)
}
