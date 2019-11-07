package main

import "fmt"

func main() {

	const a = 1 //定义常量
	const b = "/home/xiangxh/go1.13.4.linux-amd64.tar.gz"

	const ( //批量声明常量
		c, d, e = 11, 22, 33
		f, g, h = 44, 55, 66
	)

	const i, j, k = 100, 103, 300 //批量声常量

	//iota, 只能在const内部使用

	const t = iota

	const (
		l = iota
		m
		n
	)

	fmt.Println(a, b)
	fmt.Println(c, d, e, f, g, h)

	fmt.Println(i, j, k)
	fmt.Println(l, m, n)

}
