package main

import "fmt"

// 标准声明变量
//var 变量名   变量类型
//变量声明以关键字 var 开头，变量类型放在变量名后面，行尾无需分号
// 批量变量声明
/*var (
	变量名1 类型1
	变量名2 类型2
	变量名3 类型3
)
*/
var name int
var age int
var isOk bool

var (
	b string
	c float64
	d float32
	f bool
)

var name1 string = "隔壁老王"
var age1 int = 19

// 声明的变量，必须要使用，否则会报错

func foo() (int, string) {
	return 666, "Golang"
}

func main() {

	//25个关键字，37个保留字
	fmt.Println(name, age, isOk)
	fmt.Println(b, c, d, f)
	fmt.Println(name1, age1)
	//短变量声明，只能在函数内部声明
	m := 10
	fmt.Println(m)

	// 匿名变量
	x, _ := foo()
	_, y := foo()
	fmt.Println("x = ", x)
	fmt.Println("y = ", y)

	// 常量定义了，必须赋值
	// 常量
	const pi = 3.1415
	const e = "/home/xiangxh/my.cnf"

	//多常量声明
	const (
		n1 = 10
		n2 //默认和上面值一样
		n3
	)
	fmt.Println(pi, e, n1, n2, n3)
	// iota 只能使用在常量里，遇到const关键字，重置为零
	// 几个常见的iota示例
	const (
		nn1 = iota //value: 0
		nn2 = iota //value: 1
		nn3 = iota //value: 2
		nn4 = iota //value: 3
	)
	const (
		m1 = iota //value: 0
		m2        //value: 1
		m3        //value: 2
		m4        //value: 3
	)
	// 使用 _ 跳过某些值
	const (
		y1 = iota //value: 0
		y2        //value: 1
		_
		y4 //value: 3
	)
	//声明中间插队
	const (
		x1 = iota //value:0
		x2 = 100  //value:100
		x3 = iota //value:2
		x4        //value:3
	)
	const x5 = iota //value:0
	//定义数量级
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)
	//多个 iota定义在一行
	const (
		a1, b1 = iota + 1, iota + 2 //iota=0,1,2
		c1, d1
		e1, f1
	)
	fmt.Println(nn1, nn2, nn3, nn4)
	fmt.Printf("m1 = %d m2 = %d m3 = %d m4 = %d\n", m1, m2, m3, m4)
	fmt.Printf("y1 = %d y2 = %d y4 = %d\n", y1, y2, y4)
	fmt.Printf("x1 = %d x2 = %d x3 = %d x4 = %d\n", x1, x2, x3, x4)
	fmt.Printf("KB = %d MB = %d  GB = %d TB = %d PB = %d\n", KB, MB, GB, TB, PB)
}
