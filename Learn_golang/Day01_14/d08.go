package main

import (
	"fmt"
	"math"
	"strings"
)

// www.liwenzhou.com/posts/Go/02_datatype/
// go语言基本数据类型
func main() {
	//整型分为以下两个大类： 按长度分为：int8、int16、int32、int64
	//对应的无符号整型：uint8、uint16、uint32、uint64
	//其中，uint8就是我们熟知的byte型，int16对应C语言中的short型，
	//int64对应C语言中的long型
	// 十进制打印为二进制
	n := 10
	fmt.Printf("%b\n", n)
	fmt.Printf("%d\n", n)
	// 八进制
	m := 0.75
	fmt.Printf("%d\n", m)
	fmt.Printf("%o\n", m)
	// 十六进制
	f := 0xff
	fmt.Printf("%x\n", f)

	// uint8 0~255
	var age uint8
	age = 255
	fmt.Println(age)
	// 浮点数
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
	//复数 complex64 complex128

	// 布尔值
	//	Go语言中以bool类型进行声明布尔型数据，布尔型数据只有true（真）和
	//	false（假）两个值
	//注意：

	//    布尔类型变量的默认值为false。
	//    Go 语言中不允许将整型强制转换为布尔型.
	//    布尔型无法参与数值运算，也无法与其他类型进行转换。
	var b bool

	fmt.Println(b)
	b = true
	fmt.Println(b)
	//字符串
	s1 := "hello golang"
	s2 := "人生苦短，Let's go"
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println("\t制表符\n换行符")
	s3 := `
	多行字符串
	两个反引号之间的内容
	会
	原样输出
	\t
	\n
	`
	fmt.Println(s3)

	// 字符串的长度
	str := "hello_Golang"
	fmt.Println(len(str))
	str2 := "hello_谷歌"
	fmt.Println(str2)
	// 字符串拼接
	fmt.Println(str + str2)
	//	str3 := fmt.Printf("%s - %s", str, str2)
	//	fmt.Println(str3)
	// 字符串的分割 import strings
	str4 := "how are you"
	fmt.Println(strings.Split(str4, " "))
	//判断是否包含
	fmt.Println(strings.Contains(str4, "are"))
	//判断前缀
	fmt.Println(strings.HasPrefix(str4, "how"))
	// 判断后缀
	fmt.Println(strings.HasSuffix(str4, "you"))
	// 判断字符串的位置
	fmt.Println(strings.Index(str4, "you"))
	// 判断字符串出现的位置
	fmt.Println(strings.LastIndex(str4, "u"))
	// join
	str5 := []string{"what's your name?"}
	fmt.Println("[]sting", str5)
	fmt.Println(strings.Join(str5, "+"))
	// byte和rune类型
	// 字符
	// byte uint8的别名  ASCII码
	// rune int32的别名
	var c1 byte = 'c'
	var c2 rune = 'c'

	fmt.Println(c1, c2)
	fmt.Printf("c1: %T, c2: %T\n", c1, c2)

	w := "你hao"
	for i := 0; i < len(w); i++ {
		fmt.Printf("%c\n", w[i])
	}
	for _, r := range w { //会按照字符循环打印出来
		fmt.Printf("%c\n", r)
	}

}
