package main

// go语言中没有像类的概念，但是可以通过结构体struct实现oop(面向对象编程）
// struct的成员(也叫属性或字段)可以是任何类型，如普通类型、复合类型、函数、map、interface、struct等，所以我们可以理解为go语言中的“类”

type Student struct {
	name  string
	age   int
	print int
}

func StructTest() {
	//stu1 *Student = &Student{name="town",age="29",print=func(){fmt.Println(name,age)}}
	//stu1.print()
}
