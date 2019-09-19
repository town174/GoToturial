package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

//FileTest
func FileTest() {}

//ReadCharTest
func ReadCharTest() {
	file, err := os.Open("./file.txt")
	if err != nil {
		fmt.Println("open file error: ", err)
		return
	}
	defer file.Close()
	context := make([]byte, 100) //create byte array
	for {
		readNum, err := file.Read(context)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == readNum {
			break //return when end of read
		}
	}
	for k, v := range context {
		println(k, v)
	}
}

//ReadLineTest
func ReadLineTest() {
	file, err := os.Open("./file.txt")
	if err != nil {
		fmt.Println("open file error: ", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		var line []byte
		data, prefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		line = append(line, data...)
		if !prefix {
			fmt.Printf("data:%s\n", string(line))
		}
	}
}

//GetArgs
func GetArgs() {
	if len(os.Args) <= 2 {
		fmt.Println("no args")
		return
	}
	println("script name： ", os.Args[0])
	for i := range os.Args {
		fmt.Printf("this is %d arg : %s\n", i, os.Args[i])
	}
}

//flag包提供了一系列解析命令行参数的功能接口，其定义的命令行参数方式有以下几种：
//-flag //只支持bool类型
//-flag=x
//-flag x //只支持非bool类型 特别说明一个-和 -- 效果是一样的

//FlagTest
func FlagTest() {
	ip := flag.String("ip", "10.0.0.230", "server listen ip") //参数一为命令行接受的参数名称，参数二为默认值，参数三为描述信息
	port := flag.Int("port", 80, "server port")
	flag.Parse()
	fmt.Println("ip", *ip)
	fmt.Println("port", *port)
}

//多种序列化方式
//方式		优点					 缺点
//binary	性能高				     不支持不确定大小类型 int、slice、string
//json		支持多种类型			 性能低于 binary 和 protobuf
//protobuf	支持多种类型，性能高	  需要单独存放结构，如果结构变动需要重新生成 .pb.go 文件
//gob		支持多种类型			 性能低

type Student struct {
	Name  string `json:"name"` //序列化时将字段变为小写
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

//JsonSerializeTest
func JsonSerializeTest() {
	stu1 := &Student{Name: "wd", Age: 22, Score: 100}
	res, err := json.Marshal(stu1)
	if err != nil {
		fmt.Println("json encode error")
		return
	}
	fmt.Printf("json string: %s \n", res)
}

//JsonTest
func JsonDeSerializeTest() {
	date := `{"Name": "wd", "Age": 22, "Score": 100}`
	var stu1 Student
	err := json.Unmarshal([]byte(date), &stu1)
	if err != nil {
		fmt.Println("json decode error")
		return
	}
	fmt.Printf("struct obj is %s", stu1.Name)
}
