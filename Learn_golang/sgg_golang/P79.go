package main

import "fmt"

// 流程控制
// if
func main() {
	var age int
	fmt.Println("请输入你的年龄：")
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("你已经年18周岁了，已经是成人了")
	} else {
		fmt.Println("你还是未成年人，请离开。")
	}

	//golang的if还有一个强大的地方，就是条件判断语句里面允许声明一个变量，
	//这个变量的作用域只在该条件语句内起作用，条件语句外就不起作用了

	if language := "Golang"; language == "Golang" {
		fmt.Printf("Welcome to learn %s\n", language)
	}

	//fmt.Println(language) // error   undefined: language
	var x int = 4
	var y int = 1
	if x > y { //if (x >y) 官方不推荐这样写
		if y > 2 {
			fmt.Println(x + y)
		}
		fmt.Println("xxxxxxx ")
	} else {
		fmt.Printf("x = %d\n", x)
	}

	var n1 int = 10
	var n2 int = 50
	if n1+n2 > 50 {
		fmt.Printf("n1 + n2 = %d\n", (n1 + n2))
	}

	var n3 float64 = 11.0
	var n4 float64 = 17.0
	if n3 > 10.0 && n4 < 20.0 {
		fmt.Printf("和 = %f\n", (n3 + n4))
	}

	//定义两个变量int32 判断二者的和，是否被3又能被 5 整除。
	var n5 int32 = 10
	var n6 int32 = 5
	if (n5+n6)%3 == 0 && (n5+n6)%5 == 0 {
		fmt.Println("能被3又能被5整除")
	}
	// 判断一个年份是否闰年，润年的条件是符合下面二者之一：
	// (1) 年份能被4整除，但不能被100整除
	// (2) 能不400整除
	var year int32
	fmt.Printf("请输入年份：")
	fmt.Scanln(&year)
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Printf("%d 是闰年\n", year)
	} else {
		fmt.Printf("%d 不是闰年\n", year)
	}

}

