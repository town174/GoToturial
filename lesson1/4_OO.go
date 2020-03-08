package main

// go语言中没有像类的概念，但是可以通过结构体struct实现oop(面向对象编程）
// struct的成员(也叫属性或字段)可以是任何类型，如普通类型、复合类型、函数、map、interface、struct等，所以我们可以理解为go语言中的“类”

import "fmt"

//字段后接tag,表示对字段的注释，或者序列化
type Student struct {
	name string "the name of student"
	age  int    "the age of student"
	num  string "the num of student"
}

type Teacher struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//通过匿名成员可实现继承
//同一种类型匿名成员只允许最多存在一个。
//当匿名成员是结构体时，且两个结构体中都存在相同字段时，优先选择最近的字段。
type Master struct {
	Scope int
	Title string
	Teacher
}

func Inherit() {
	mst := &Master{Scope: 19, Title: "the great"}
	mst.Teacher.Name = "Lee HH"
	mst.Teacher.Age = 60
	fmt.Println(mst.Teacher.Name, mst.Teacher.Age, mst.Title)
}

func StuConstruct(nameA string, numA string, ageA int) *Student {
	return &Student{name: nameA, age: ageA, num: numA}
}

func StructTest() {
	//stu1 *Student = &Student{name="town",age="29",print=func(){fmt.Println(name,age)}}
	//stu1.print()
	var stu1 Student
	stu1.age = 22
	stu1.name = "town"
	stu1.num = "090103041012"
	fmt.Println(stu1.name, stu1.num, stu1.age)

	var stu2 = new(Student)
	stu2.name = "abby"
	stu2.age = 22
	fmt.Println(stu2.name, stu2.num, stu2.age)

	var stu3 *Student = &Student{name: "rose", age: 18, num: "class3"}
	fmt.Println(stu3.name, (*stu3).name) //rose  rose
}
