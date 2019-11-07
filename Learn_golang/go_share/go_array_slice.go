package main

// http://zengzhihai.com/study/wiki/type/Z29fc3R1ZHlfaW5mbw==/id/59ac990ca55d5081678b456a
import "fmt"

func main() {
	// 数组的定义
	// 数组是go语言中最常用的数据结构之一，数组就是指一系列同一类类型数据集合。数组中包含的每个数据被称为数组元素
	// 一个数组包含的元素个数被称作为数组的长度。是值类型。
	/* 数组有3 中创建方式：
	[length]Type

	[N]Type{v1,v2,...vn}

	[...]Type{v1,v2,....vn}
	*/
	arr1 := [5]int{1, 2, 3, 4, 5}         //创建数组大小为5
	arr2 := [5]int{1, 2}                  //创建数组大小为5，但是内容没有写的用0代替
	arr3 := [...]int{11, 22, 33, 44, 55}  //数组未定义长度
	arr4 := [5]int{1: 102, 2: 300, 3: 34} //数组有 key  value
	arr5 := [...]int{22: 33, 44: 55}      //数组长度未定义，并且是key value形式
	arr6 := [...]string{"Golang", "Python", "Shell", "Perl"}
	arr5[1] = 2400
	// arr5[6] = 223 // 数组不支持最大的key进行赋值
	arr := [...]int{110, 120, 130, 140, 150, 160}
	fmt.Println("arr = ", arr)
	fmt.Println("arr1 = ", arr1)
	fmt.Println("arr2 = ", arr2)
	fmt.Println("arr3 = ", arr3)
	fmt.Println("arr4 = ", arr4)
	fmt.Println("arr5 = ", arr5)
	fmt.Println("arr6 = ", arr6)

	/*
		切片的定义
		go语言中，切片是长度可变，容量固定的相同的元素序列，go语言的切片本质是一个数组，容量固定是
		为数组的长度固定的，切片的容量即隐藏在数组的长度。长度可变指的是数组长度的范围内可变
		go语言提供了数组切片，（slice)弥补数组的不足，其实数组切片就是一个指向数组的指针，实际上它
		拥有自己的数据结构，而不仅仅是个指针。数组切片的数据结构可以抽象以下3个变量：
		1、一个指向原生数组的指针
		2、数组切片中的元素个数
		3、数组切片一分配的储存空间
		切片的创建方式有以下4种：
		1、make([]Type,length,capacity)
		2、make([]Type,length)
		3、[]Type{}
		4、[]Type{v1,v2,...vn}
	*/

	s := []int{1, 2, 3} //直接初始化切片，[]表示是切片类型，{1,2,3}表示初始化值依次是1,2,3
	// 其中 cap=len=3
	fmt.Println("s = ", s)

}
