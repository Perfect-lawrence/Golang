package main

import "fmt"

// 函数
func sayHello() {
	fmt.Println("Go语言中的函数") //定义一个不需要参数也没有返回值的函数
}

// 带参数的函数
func argFunc(name string) {
	fmt.Println("函数的参数是: ", name)
}

// 定义接收多个参数的函数并且又一个返回值

func intSum(a int, b int) int {
	res := a + b
	return res

}

// 方法2 返回值简写    返回值定义一个 别名： result
func intSum2(a, b int) (result int) {
	result = a + b
	return
}

// 函数接收可变参数，在参数后面加上 ... 表示可变参数
// 可变参数在函数体中是切片类型
// 固定参数和可变参数同时出现是，可变参数要放在最后
func intSum3(a ...int) int {
	r := 0
	for _, arg := range a {
		r = r + arg
	}
	return r
}

//// 固定参数和可变参数同时出现是，可变参数要放在最后
// go 语言没有默认参数
func intSum4(a int, b ...int) int {
	result := a
	for _, arg := range b {
		result = result + arg
	}
	return result
}

//go 语言中函数参数类型简写

// 定义具有多个返回值的函数,多返回值也支持简写
func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

// defer 语句
/*

Go语言中的defer语句会将其后面跟随的语句进行延迟处理。
在defer归属的函数即将返回时，
将延迟处理的语句按defer定义的逆序进行执行，也就是说
，先被defer的语句最后被执行
，最后被defer的语句，最先被执行。
*/

// defer 延迟执行
func deferFunc() {
	fmt.Println("Start .....")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	defer fmt.Println(4)
	fmt.Println("End......")
}

// 函数进阶之变量作用域
// 定义一个全局变量
var num = 10
var i = "外层变量 i"

// 定义函数
func testGolbal() {
	// 可以在函数中访问全局变量
	num := 100
	fmt.Println("num 变量：", num)
	// 1 现在自己函数中找，找到了就用自己函数中的
	// 2 函数中找不到变量就往外层找全局变量
	fmt.Println(num)
	// 变量i此时只在for循环的语句中生效
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println(i) //外层访问不了到内部for语句块中的变量
}

func add(x, y int) int {
	return x + y
}

func sub2(x, y int) int {
	return x - y
}

// 函数作为参数传入

func calc2(x, y int, op func(int, int) int) int {
	return op(x, y)
}
func main() {
	// 函数的调用
	sayHello()
	argFunc("lawrence")
	intSum(23, 35)
	manyArg := intSum(3, 5)
	fmt.Println("定义接收多个参数的函数并且又一个返回值: ", manyArg)
	rr := intSum2(12, 34)
	fmt.Println("给函数返回值定义一个 别名：", rr)
	s1 := intSum3()
	s2 := intSum3(1)
	s3 := intSum3(1, 2)
	s4 := intSum3(3, 4, 5)
	s5 := intSum3(6, 7, 8, 9)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)

	s6 := intSum4(10)            // a=10,b=[]
	s7 := intSum4(11, 12)        // a = 11 b=[12]
	s8 := intSum4(100, 110, 120) // a=100 b = [110,120]
	fmt.Println("intSum4: ", s6)
	fmt.Println("intSum4: ", s7)
	fmt.Println("intSum4 :", s8)

	x, y := calc(100, 200)
	fmt.Printf("x = %d  y = %d\n", x, y)
	//由于defer语句延迟调用的特性，所以defer语句能非常方便的处理资源释放问题。
	//比如：资源清理、文件关闭、解锁及记录时间等
	deferFunc()
	testGolbal()
	fmt.Println("#######函数可以作为变量 testGlobal  Start ")
	abc := testGolbal // 函数可以作为变量
	fmt.Printf("%T\n", abc)
	fmt.Println("函数可以作为变量 testGloba  End")

	abc()

	fmt.Println("函数作为参数")
	qq := calc2(12, 22, add) // qq 接受 calc2
	fmt.Println("函数作为参数:", qq)

	qq2 := calc2(20, 23, sub2) // qq2 接受 calc2
	fmt.Println(qq2)

}
