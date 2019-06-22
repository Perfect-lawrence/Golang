package main

import "fmt"

// www.liwenzhou.com/posts/Go/04_basic
// 流程控制 if  for  switch goto
func main() {
	var score = 65
	// 基本写法
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("c")
	}
	// if判断的特殊写法

	if s := 70; s >= 90 {
		fmt.Println("AA")
	} else if s > 75 {
		fmt.Println("BB")
	} else {
		fmt.Println("cc")
	}

	// for循环
	/* for 初始语句;条件表达式;结束语句{
	    循环体语句
	} */

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//省略初始语句，但是必须保留初始语句后面的分号
	w := 0
	for ; w < 3; w++ {
		fmt.Println(w)
	}

	//省略初始语句和结束语句
	var j = 2
	for j > 0 {
		fmt.Println(j)
		j--
	}

	// 无限循环
	//	for {
	//		fmt.Println("hello golang")
	//	}
	// break
	for z := 0; z < 5; z++ {
		if z == 3 {
			break
		}
	}
	/*	for range(键值循环)
		Go语言中可以使用for range遍历数组、切片、字符串、map 及通道（channel）。 通过for range遍历的返回值有以下规律：

			    数组、切片、字符串返回索引和值。
			    map返回键和值。
			    通道（channel）只返回通道内的值。
	*/

	// switch case
	score2 := 80
	switch score2 {
	case 90:
		fmt.Println("score2 = 90")
	case 80:
		fmt.Println("score2 = 80")
	case 70:
		fmt.Println("score2 = 70")
	case 60:
		fmt.Println("score2 = 60")
	default:
		fmt.Println("score2 < 60")
	}

	// case 一次判断多个值
	num := 5
	switch num {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")
	}
	// case 使用表达式
	age := 30
	switch {
	case age > 18:
		fmt.Println("你已经年满18周岁了")
	case age < 18:
		fmt.Println("你未满18周岁")
	default:
		fmt.Println("123，456")
	}
}
