package main

//数组相关内容

import (
	"fmt"
)

func main() {
	var a [3]int
	var b [4]int
	fmt.Println(a)
	fmt.Println(b)
	// 数组的初始化  数组的长度一定要时常量，可以确定长度，不可以定义之后还能修改数组的长度
	//1 定义时使用初始值列表的方法初始化
	var cityArray = [5]string{"湖南", "广东", "浙江", "四川", "广西"}
	fmt.Println(cityArray)
	// 2 编译器推导数组长度
	var boolArray = [...]bool{true, false, false, true, false}
	fmt.Println(boolArray)
	// 数组可以通索引访问
	fmt.Println(boolArray[3])
	// 3 使用索引值方式初始化
	var langArray = [...]string{1: "Golang", 2: "Python", 3: "java", 7: "PHP"}
	fmt.Println(langArray)
	fmt.Println("langArray[7]:", langArray[7])
	fmt.Printf("%T\n", langArray) // 数组长度是 8

	// 数组的遍历
	// 方法1
	var strArray = [...]string{"pgsql", "mysql", "mariadb", "percona_mysql"}
	for i := 0; i < len(strArray); i++ {
		fmt.Println(strArray[i])

	}
	// 方法 2
	for _, v := range strArray {
		fmt.Println("方法2", v)
	}

	// 二维数组
	nosqlArray := [4][2]string{
		{"redis", "mongodb"},
		{"memcache", "Hbase"},
		{"CouchDB", "Etcd"},
		{"Neo4J", "InfoGrid"},
	}
	fmt.Println(nosqlArray) // 数组是从0开始计算的
	fmt.Println(nosqlArray[2][1])

	// 二维数组遍历
	// 注意：多维数组 zArray := [...][2] // ok
	// zArray := [...][...]  //error
	for _, v1 := range nosqlArray {
		//		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	// 数组是值类型，实际是一个拷贝
	// 二维数组
	x := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(x)
	//当作参数传到函数里，里面的 x 值还是没有变
	f1(x)
	fmt.Println("传入函数里x值都没有变，所以说数组是值类型，实际是一个拷贝：", x)
	y := x
	fmt.Println("y := x 拷贝x值", y)
	// 求数组{1, 3, 5, 7, 8}所有元素的和
	var sumArray = [...]int{1, 3, 5, 7, 8}
	var s int
	for _, v := range sumArray {
		s += v
		fmt.Println(s)
	}

}
func f1(a [3][2]int) {
	a[0][0] = 100
}
