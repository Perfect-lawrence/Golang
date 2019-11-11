package main

import "fmt"

// 声明变量
// 变量必须声明后  才能使用
//go 语言中推荐使用   驼峰命名
// var name string
// var age int
// var isOk bool

var (
	name string // ""
	age  int    // 0
	isOk bool   // false
)

func main() {
	name = "Welcome to golang"
	age = 16
	isOk = true
	// go 语言中变量声明必须使用，否则会编译不过
	fmt.Printf("name = %s\n", name)
	fmt.Println("age = ", age)
	fmt.Printf("isOk = %v\n", isOk)

	//声明变量同时赋值
	var s1 string = "world"
	fmt.Println("s1 = ", s1)

	// 类型推倒(根据值判断该变量是什么类型)

	var s2 = "200"
	fmt.Println("s2 = ", s2)

	// 短变量声明，只能在函数内使用
	s3 := "短变量声明"
	fmt.Println("s3 = ", s3)

	// 匿名名变量   使用  _(下划线表示匿命名变量)

}
