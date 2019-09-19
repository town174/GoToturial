package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"reflect"
)

//GetTest
func GetTest() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	headers := resp.Header
	for k, v := range headers {
		fmt.Printf("k=%v, v=%v \n", k, v)
	}
	fmt.Printf("resp status %s,statusCode %d\n", resp.Status, resp.StatusCode)
	fmt.Printf("resp Proto %s\n", resp.Proto)
	fmt.Printf("resp transfer encoding %v\n", resp.TransferEncoding)
	fmt.Printf("resp Uncompressed %t\n", resp.Uncompressed)

	fmt.Println(reflect.TypeOf(resp.Body))
	buf := bytes.NewBuffer(make([]byte, 0, 512))
	length, _ := buf.ReadFrom(resp.Body)
	fmt.Println(len(buf.Bytes()))
	fmt.Println(length)
	//fmt.Println(string(buf.Bytes()))
}
