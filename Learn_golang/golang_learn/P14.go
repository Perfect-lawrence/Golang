package main

import (
	"fmt"
	_ "os"
	"strings"
)

//go语言字符串是用双引号包裹的！！！
// go 语言单引号 是 字符
/*
字节： 1 字节 = 8Bit（8个二进制位)
1 字符 'A' = 1 个字节
1 个utf8编码
*/
func main() {
	s := "go语言字符串是用双引号包裹的"
	c1 := 'h'
	c2 := '1'
	c3 := '海'

	fmt.Println("s = ", s)
	fmt.Printf("c1 = %v c1 = %v c3 = %v\n", c1, c2, c3)

	//转义 符 \
	pathFile := "F:\\Golang_code\\src\\b1024.cn\\studygo\\day01\\08string"
	fmt.Println(pathFile)

	// 多行 的字符串
	s1 := `
	this is my teacher
	is that ok
	i am fine
	`
	T := `F:\\Golang_code\\src\\b1024.cn\\studygo\\day01\\08string`

	fmt.Println(s1)
	fmt.Println(T)

	//字符串拼接
	name := "hello world"
	name2 := "welcome to golang world"
	ss := name + name2
	fmt.Println(ss)

	ss1 := fmt.Sprintf("%s      %s", name, name2)
	fmt.Println(ss1)

	// 分隔
	ret := strings.Split(T, "\\")
	fmt.Println(ret)

	//包含
	fmt.Println(strings.Contains(ss, "golang"))
	// 前缀
	fmt.Println(strings.HasPrefix(name2, "welcome"))

	//后缀
	fmt.Println(strings.HasSuffix(name, "world"))

}
