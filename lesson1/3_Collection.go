package main

import (
	"fmt"
	"sort"
)

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

//数组测试，数组需要定长
func ArrayTest() {
	//var a1 = [2]int
	//a1[0] = 1
	a2 := make([]int, 3)
	a3 := [2]int{1, 2}
	a4 := [...]int{1, 2}
	a5 := [...]int{0: 8, 1: 1, 3: 3}
	//println("a1 [2]int len: %d", len(a1))
	fmt.Printf("a2 make([]int,3) len: %d \n", len(a2))
	fmt.Printf("a3 [2]int{1,2} len: %d \n", len(a3))
	fmt.Printf("a4 [...]int{1,2} len: %d \n", len(a4))
	fmt.Printf("a5 [...]int{1:1,2:2} len: %d \n", len(a5))
	for i := 0; i < len(a5); i++ {
		fmt.Println(a5[i])
	}
	ArrayAssignment(&a2)
	fmt.Println(a2[0])
}

func ArrayAssignment(a *[]int) {
	(*a)[0] = 100
	return
}

//切片测试，切片的长度是可变的
//切片特性和数组一样（遍历  计算长度）
func SliceTest() {
	slice1 := make([]int, 10, 10)
	slice2 := []int{2, 2, 3}
	fmt.Printf("slice1 [2]int{1,2} len: %d \n", len(slice1))
	slice2 = append(slice2, 3)
	//slice2 = append(slice2, slice1)
	fmt.Printf("silce2 []int{1, 2} len: %d  cap：%d \n", len(slice2), cap(slice2))
}

//map测试
func MapTest() {
	//map属于引用类型，声明是不会分配内存的，需要make初始化分配内存
	var m1 map[string]string
	m1 = make(map[string]string, 8)
	fmt.Printf("m1 len: %d \n", len(m1))
	m1["k1"] = "v1"
	fmt.Println("m1[k1] value:", m1["k1"])
	//map可以嵌套
	m2 := make(map[string]map[string]string, 10)
	m2["k2"] = make(map[string]string, 2)
	m2["k2"]["kk1"] = "vv1"
	m2["k2"]["kk2"] = "vv2"
	fmt.Println(m2)
	//删除
	delete(m1, "k1")
	fmt.Println("m1[k1] value:", m1["k1"])
	//遍历
	m3 := map[string]string{"name": "wd", "age": "22"}
	for k, v := range m3 {
		fmt.Println(k, v)
	}
	//排序  map目前无序的，并且无内置排序方法
	m4 := map[string]string{"4": "d", "1": "a", "3": "c", "2": "b"}
	var keys []string
	for k := range m4 {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Print(keys);
}
