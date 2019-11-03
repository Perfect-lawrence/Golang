package main

import (
	"fmt"
	"strconv"
)

//golang 中string转成基本数据类型
func main() {
	var str string = "true"
	var b bool
	// b,_ = strconv.ParseBool(str)
	// 说明：
	// 1、strconv.ParseBool(str) 函数会返回两个值（value bool,err error)
	// 2、 因为我只想获到 value bool, 不想获取 err 所以使用_忽略
	b, _ = strconv.ParseBool(str)
	fmt.Printf("str to b bool type is = %T b = %v\n", b, b)

	var str2 string = "3490205496"
	var nn int64
	nn, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("str2 to nn type is %T nn = %v\n", nn, nn)

	var nn2 int
	nn2 = int(nn)
	fmt.Printf("nn to nn2 type is %T nn2 = %v\n", nn2, nn2)

	var str3 string = "134.35456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("str3 to f1 Type is %T f1 = %v\n", f1, f1)

}

