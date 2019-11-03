package main

import (
	"fmt"
)

//golang中的基本数据类型的转换
func main() {
	var t int32 = 100
	// 希望将 t => float
	var t2 float32 = float32(t)
	var t3 int8 = int8(t)
	var t4 int64 = int64(t)
	fmt.Printf("t = %v t2 = %v t3 = %v t4 = %v\n", t, t2, t3, t4)
	//被转换的是变量存储的数据（即值），变量本身的数据类型并没有变化
	fmt.Printf("i type is %T\n", t) //int32

	var num int64 = 888888
	var num2 int8 = int8(num)
	fmt.Println("num2 = ", num2)

	var e float64
	var e2 int8
	e = 12.0
	e2 = int8(e) + 3
	fmt.Println(e, e2)

}

