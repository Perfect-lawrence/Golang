package main

import "fmt"

// 常量  const
// 在程序运行期间不会改变的量
const pi = 3.1415
const e = 2.7182

//批量声明常量
const (
	statusCode = 200
	notFound   = 404
)

const (
	n1 = 200
	n2
	n3
)

//iota  类似枚举
const (
	a1 = iota //0
	a2 = iota //1
	a3        // 2
)

const (
	b1 = iota // 0
	b2        //1
	_
	b3 //3
)

// 插队
const (
	c1 = iota
	c2 = 100
	c3 = iota
	c4
)

// 多个常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2 // d1: 1  d2: 3
	d3, d4 = iota + 1, iota + 2 // d3: 2   d4: 3
)

//定义数量级 10 次方
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)

	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)

	fmt.Println("b1: ", b1)
	fmt.Println("b2: ", b2)
	fmt.Println("b3: ", b3)

	fmt.Println("c1: ", c1)
	fmt.Println("c2: ", c2)
	fmt.Println("c3: ", c3)
	fmt.Println("c4: ", c4)

	fmt.Println("d1: ", d1)
	fmt.Println("d2: ", d2)
	fmt.Println("d3: ", d3)
	fmt.Println("d4: ", d4)

	fmt.Println("KB = ", KB)
	fmt.Println("MB = ", MB)

	fmt.Println("GB = ", GB)
	fmt.Println("TB = ", TB)
	fmt.Println("PB = ", PB)

}
