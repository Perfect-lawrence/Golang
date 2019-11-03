package main

import (
	"fmt"
	_ "os"
	"strconv"
)

//基本数据类型转string
func main() {
	var n1 int = 99
	var n2 float64 = 24.567
	var b bool = true
	var myChar byte = 'h'
	var str string // 空的str
	// 第一种方式来转换  fmt.Sprintf
	str = fmt.Sprintf("%d", n1)
	fmt.Printf("str type is = %T str = %v\n", str, str)

	str = fmt.Sprintf("%f", n2)
	fmt.Printf("n2 to str type is = %T  str = %q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("b to str type is = %T str = %q\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("myChar to str type is = %T str = %q\n", str, str)
	// 第二种 方式 strconv 函数
	var ss int = 100
	var ss2 float64 = 25.678
	var b2 bool = true
	str = strconv.FormatInt(int64(ss), 10)
	fmt.Printf("ss to str type is = %T str = %q\n", str, str)
	// 说明：'f' 格式，10： 表示小数保留10位，64:小数是64位 float64
	str = strconv.FormatFloat(ss2, 'f', 10, 64)
	fmt.Printf("ss2 to str type is = %T str = %q\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("b2 to str type is = %T str = %q\n", str, str)

	// strconv包中有一个函数 Itoa
	var n5 int64 = 45697
	str = strconv.Itoa(int(n5))
	fmt.Printf("n5 to str type is = %T str = %q\n", str, str)

}

