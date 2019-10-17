package main

import (
	"fmt"
)

// 指针 3 个概念： 指针地址，指针类型  指针值
// 指针两个符号： & (取地址) 和 * (根据地址取值)
// 指针地址 和指针类型
//go 语言中的值类型(int float bool string array struct)

// 指针
func zhiZhen() {
	a := 10 //int 类型
	b := &a // *int 类型(int指针)
	fmt.Printf("%v %p\n", a, &a)
	fmt.Printf("%v %p\n", b, b)
	// 变量b是一个int类型的指针(*int) ,它存储的是变量a的内存地址
	c := *b
	fmt.Println(c) // 10
}

//指针传值操作
func modify1(x int) {
	x = 100
}
func modify2(y *int) {
	*y = 100 // 整形的指针
}

//func make(t Type, size ... InterType) type
// new  and   make
//new函数不太常用，使用new函数得到的是一个类型的指针，并且该指针
// 对应的值为该类型的零值(nil).

//new和make的区别
// 1 二者都是用来做内存分配的
// 2 make只用于slice  map  channel 的初始化，返回的还是 三个
// 引用类本身
// 3 而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针

func nn() {
	a := new(int)
	//*a
	b := new(bool)
	fmt.Printf("%T\n", a) // *int
	fmt.Printf("%T\n", b) // *bool
	fmt.Println(*a)
	fmt.Println(*b)
}

func main() { //指针
	zhiZhen()
	a := 1
	modify1(a)
	fmt.Println(a) //1
	modify2(&a)
	fmt.Println(a)

	nn()
}
