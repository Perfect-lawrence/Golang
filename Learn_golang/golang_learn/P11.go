package main

import "fmt"

func main() {
	// 十进制
	var i1 = 101
	fmt.Printf("十进制: %d\n", i1)
	fmt.Printf("把十进制数转换成二进制: %b\n", i1)
	fmt.Printf("把十进制数转换成八进制: %o\n", i1)
	fmt.Printf("把十进制数转换成十六进制: %x\n", i1)

	// 八进制
	i2 := 077
	fmt.Printf("八进制: %d\n", i2)

	// 十六进制
	i3 := 0x1234567
	fmt.Printf("十六进制: %d\n", i3)

	// 查看变量的类型
	fmt.Printf("%T\n", i3)

	// 声明int8类型的变量
	i4 := int8(9)
	fmt.Printf("%v\n", i4)
}

