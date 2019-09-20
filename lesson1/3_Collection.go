package main

//内置函数
//make :用来分配内存，主要分配引用类型，返回Type本身(只能应用于slice, map, channel)
//new ：分配内存，主要分配值类型，返回指向Type的指针，如int
//cap  ：容量，容积capacity
//copy ：复制slice，返回复制的数目
//len ：返回长度

//append ：追加元素到slice里,返回修改后的slice
//close ：关闭channel
//delete ：从map中删除key对应的value
//panic  ： 用于异常处理，停止常规的goroutine
//recover ：用于异常处理，允许程序定义goroutine的panic动作

func ArrayTest() {
	a1 := []int
	
}
