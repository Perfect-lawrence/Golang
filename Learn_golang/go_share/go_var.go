package main

import "fmt"

func main() {
	var a int      //声明一个int类型的变量
	var b struct { //声明一个结构体
		name string
	}
	var c = 8 // 声明变量同时赋值
	var (     //批量声明变量
		d int
		e string
	)
	var name1 int = 5 // 声明变量name1，并初始化

	var f, g int //一行声明多个变量

	//同一行初始化多个变量，不同类型也可以，这里面默认初始化值，根据值进行定义了类型
	var h, i, j = 23, "Golang", 0.56

	name2 := "Hello, Welcome to Golang World!"

	fmt.Println(a, b, c, d, e, f, g, h, i, j)
	fmt.Println(name1, name2)
/*
小结：
1、变量声明(variable declaration)中的初始化表达是(initalization expressions）的求值顺序



*/

}
