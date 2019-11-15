package main

import (
	"fmt"
	_ "math"
)

// 浮点数
func main() {
	//	Max.Float32
	//fmt 占位符
	var n = 100
	// 查看类型
	fmt.Printf("%T\n", n)
	fmt.Printf("%v\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	var s = "golang 海信子"
	fmt.Printf("字符串: %s\n", s)
	fmt.Printf("字符串: %v\n", s)
	fmt.Printf("字符串: %#v\n", s)
}

