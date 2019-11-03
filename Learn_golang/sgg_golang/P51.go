package main

import (
	"fmt"
)

//复杂类型 ：指针（pointer)
func main() {
	// 基本数据类型在内存布局, 值类型（指针)
	//指针类型，指针变量存的是一个地址，这个地址指向的空间存的是才是值
	var i int = 10
	// i 的内存地址
	//a = &i
	fmt.Println("i 的内存地址：", &i)

	//下面的var ptr *int = &i
	// 1、ptr是一个指针变量
	// 2、ptr 的类型是 *int
	// 3、ptr 本身的值 &i
	var ptr *int = &i
	fmt.Printf("ptr 的内存地址 %v\n", ptr)

}

