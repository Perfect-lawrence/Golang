package main

import "fmt"

func main() {
	var str string //声明字符串

	var str1 = "Lets's get together"

	var str2 = " Just do it, you can make it"

	//字符串函数 https://studygolang.com/pkgdoc

	str = "Welcome to Golang World" //赋值
	ch := str[0:6]                  //获取第一个单词
	l := len(str)                   //字符串长度
	fmt.Println(str, ch, l)

	gogo := "lawrence is 向"

	for i := 0; i < len(gogo); i++ {

		fmt.Println(gogo[i])

	}

	for i, s := range gogo {
		fmt.Println(i, "Unicode (", s, ") string-", string(s))
	}

	r := []rune(gogo) //这是一个切片. rune是一种数据类型 rune = int32

	fmt.Println("rune = ", r)

	for i := 0; l < len(gogo); i++ {

		fmt.Println("r[", i, "]=", r[i], "string = ", string(r[i]))

		fmt.Println(str1, str2)
	}

	/*
	   小结：
	   len函数是go内?函数，不引入strings包即可使用，len(string)返回的是字符串的字节数。len函数所支持的入参类型如下：
	   len(Array) 数组的元数个数
	   len(*Array) 数组指针中的元素个数，如果入参为nil则返回0
	   len(Slice) 数组切片中元素个数，如果入参为nil则返回0
	   len(Map) 字典中元素个数，如果入参为nil则返回0
	   len(Channel) Channel buffer队列中元素个数

	*/

}
