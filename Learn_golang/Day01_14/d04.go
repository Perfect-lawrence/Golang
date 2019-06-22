package main

import "fmt"

// go fmt xx.go  //go 代码格式化
var n int

//n = 100 //error
func main() { //“{” 必须和函数在同一行

	//单行注释
	/*
		多行注释/块注释
	*/
	n = 100 //ok
	fmt.Println("hello 沙河娜扎")
}
