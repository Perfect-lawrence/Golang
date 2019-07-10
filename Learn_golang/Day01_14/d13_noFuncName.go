package main

import (
	"fmt"
	"strings"
)

// 匿名函数和闭包
// 定义一个函数它的返回值是一个函数
// 把函数作为返回值
func a() func() {
	return func() { //把函数作为返回值
		fmt.Println("闭包")
	}
}

//一个简单的 闭包
func b() func() {
	name := "闭包b函数"
	return func() {
		fmt.Println("hello 一个简单的闭包:", name)
	}
}

// 闭包 = 函数 + 外层变量的引用
func ww(name string) func() {
	return func() {
		fmt.Println("hello 常见的闭包形式：", name)
	}
}

// 使用闭包做文件后缀名检查的
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		// import "strings"
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 闭包 计算
func calc3(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

// panic and recover

func aa() {
	fmt.Println("func aa")
}
func bbb() {
	defer func() { // recover 必须配合 defer 一起用
		err := recover()
		if err != nil {
			fmt.Println("func bbbbbb error")
		}
	}() // 一定要在panic 出现的前面定义defer
	panic("panic in bbb ")
}
func cc() {
	fmt.Println("func cc")
}

func main() {
	//匿名函数
	func() {
		fmt.Println("匿名函数")
	}()
	r := a()

	r() //相当于执行 a 函数内部的匿名函数

	bb := b()
	bb() // 相当于执行 b 函数

	// 使用闭包做文件后缀名检查的
	tt := ww("Golang")
	tt()

	ee := makeSuffixFunc(".txt")
	ret := ee("闭包")
	fmt.Println(ret)

	e2 := makeSuffixFunc(".avi")
	big := e2("bigMovie")
	fmt.Println(big)

	x, y := calc3(100)
	ret1 := x(200) // base = 100 + 200
	fmt.Println(ret1)
	ret2 := y(200) // base = 300 - 200
	fmt.Println(ret2)

	// panic and recover

	aa()
	bbb()
	cc()
}
