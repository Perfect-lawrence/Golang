package main

import "fmt"

// www.liwenzhou.com/posts/Go/03_operators/
// Go 语言中的运算符
func main() {
	//1 算术运算符 + - * / %
	a := 10
	b := 20
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(5 / 2)
	fmt.Println(5 % 2)
	//关系运算符 == != > < <= >=
	a++
	fmt.Println(a)
	fmt.Println("关系运算符 == != > < <= >=")
	fmt.Println(a > b)
	fmt.Println(a != b)
	fmt.Println(a == b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)

	// 逻辑运算符 && || !
	c := 10
	d := 30
	fmt.Println("逻辑运算符 && || !")
	fmt.Println(a == c && b > d)
	fmt.Println(!(a == c))
	fmt.Println(c == d || c == a)

	//位运算符 & | ^ << >>
	aa := 1 // 001
	bb := 5 // 101
	fmt.Println("位运算符 & | ^ ")
	fmt.Println(aa & bb) //001 => 1
	fmt.Println(aa | bb) // 101 => 5
	fmt.Println(aa ^ bb) // 100 => 4

	cc := 4
	fmt.Println(aa << 2) // 左移两位   100 => 4
	fmt.Println(cc >> 2) //右移两位    1 => 1

	dd := 1
	fmt.Println(dd << 10) //1024

	//赋值运算符 =  +=  -=  *=  /=  %=  <<=  >>=  &=  |=  ^=
	var ee int
	ee = 5
	ee += 1
	fmt.Println("赋值运算符 =  +=  -=  *=  /=  %=  <<=  >>=  &=  |=  ^=")
	fmt.Println(ee)
}
