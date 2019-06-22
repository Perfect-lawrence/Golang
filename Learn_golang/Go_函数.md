# 函数
Go 语言函数定义
=============

### 函数的定义：
```markdown
func FuncName(/*参数列表*/) (r01 type1,r02 type2/*返回类型*/){
    // 函数体
    
    return v01,v02 //返回多个值
}
//函数名字首字母大写为public  函数名字首字母小写为private

```
### 无参无返回值--函数的定义

```
package main //必须

// 无参无返回值--函数的定义

import (
    "fmt"
)
func MyFunc(){
    a := 666
    fmt.Println("a =", a)
}

func main(){
    MyFunc()
}

```

## 有参无返回值--函数的定义 -- 普通参数列表

```
package main

import "fmt"
// 有参无返回值函数的定义，普通参数列表
// 定义函数时，在函数名后面（）定义的参数叫形参
// 参数传递，只能是由实参传递给形参，不能反过来，单向传递

func MyFunc01(a int){
    //a = 111
    fmt.Println( "a =", a)
}

func MyFunc02(a int, b int){
    fmt.Printf("a = %d, b = %d\n", a,b) 
}
// 相同类型数据类型
func MyFunc03(a,b int){

    fmt.Printf("a = %d, b = %d\n" ,a,b)
}

// 不同数据类型参数
func MyFunc04(a int, b string,c float64){

    fmt.Printf("a = %d, b = %s , c = %f\n " ,a,b,c)
}

func MyFunc05(a, b string,c float64,d,e int){

 //   fmt.Printf("a = %s, b = %s , c = %f, d = %d, e = %d\n " ,a,b,c,d,e)
}


func MyFunc06(a string, b string,c float64,d int,e int){

 //   fmt.Printf("a = %s, b = %s , c = %f, d = %d, e = %d\n " ,a,b,c,d,e)
}


func main(){
// 有参数无返回值函数调用： 函数名（所需参数)
// 调用函数传递的参数叫实参

    MyFunc01(666)
    MyFunc02(222,333)
    

}
```

## 函数--不定参数类型--列表

```markdown

package main

import "fmt"


func MyFunc07(a int,b int){ // 固定参数

}

// ...int 类型这样类型  ...type 不定参数类型
// 注意：不定参数  一定（只能）放在 形参中的最后一个参数

func MyFunc08(args ...int){ //传递的实参可以是0或多个
    fmt.Println("len(args) = ",len(args)) //获取用户传递参数的格数
    
fmt.Println("++++++++++++++++++++++")
    

    for i := 0; len(args); i++{
        fmt.Printf("args[%d] = %d\n", i, args[i])
      }
      

//返回2个值，第一个值是下标，第二个值是下标对应的数
    for i, data := range args { //迭代
        fmt.Printf("args[%d] = %d\n", i ,data)
    }
}

func MyFunc09(a int, args ...int){

}

func main(){
    MyFunc07(88,99)
    MyFunc08() // 0个参数
    MyFunc08(12) // 1 个参数
    MyFunc08(1,2,3) // 多个参数，不定参数
}
```
 

## 无参有返回值

```markdown
pacakge main

import "fmt"

func MyFunc01() int{
    return 123
}

// 给返回值起一个变量，go推荐写法
func MyFunc02()(rs int){
    return 888

}

// 以下范例是常用写法
func MyFunc03()(result int){
    result = 234
    return 
}
func main(){
//无参数返回值函数的调用
    var a int
    a = MyFunc01()
    
    fmt.Println("a = ",a)
    
    b := MyFunc02()
    fmt.Println("b = ",b)
    
    c := MyFunc03()
    fmt.Println("c = ",c)

}

```

## 函数有多个返回值
```markdown
package main

import "fmt"

func MyFunc01()(int,int,int){

    return 1,2,3
}

//go官方推荐写法
func MyFunc02()(a int,b int,c int){

    a,b,c = 110,120,119
    return
}

func MyFunc03()(a,b,c int){
    a,b,c = 4,5,6
    return

}
func main(){
    //函数的调用
    a, b, c := MyFunc02()
    fmt.Printf("a = %d, b = %d, c = %d\n",a,b,c)

}
```

## 有参有返回值的函数使用
```markdown
package main

import "fmt"

func MaxAndMin(a ,b int)( max ,min int{
    if a > b{
        max = a 
        min = b
    }else{
        max = b
        min = a
    }
}

func main(){

    max, min := MaxAndMin(10,20)
    fmt.Printf("max = %d, min = %d\n",max,min)
    
//通过匿名变量丢弃某个返回值
    aa, _ := MaxAndMin(100,200)
    fmt.Printf("a = %d\n",a)

}
```

## 递归函数
```markdown
package main

import "fmt"

func funcB()(b int){
    b = 222
    fmt.Println("funcB b = ",b)
    return

}

func funcA()(a int){
    a = 111
     // 调用funcB一个函数
    b := funcB()
    fmt.Println("funcB b = ",b)
    
   
    fmt.Println("funcA a = ",a)
    
    return //返回

}

// 普通函数的调用流程
func main(){
    fmt.Println("main")
    a := funcA()
    fmt.Println("main a = ",a)

}
```

## **函数递归调**用的流程
```markdown
package main

import "fmt"

func test(a int) {
     if  a == 1 {
         fmt.Println("a = ",a)
         return // 终止函数调用
    }
    // 函数调用自身
    test(a - 1)
    
    fmt.Println("a = ",a)
}

func main(){
    test(3)
    fmt.Println("main")


}
```
## 递归函数实现--数字累加
```markdown
package main

import "fmt"

// 实现 1 + 2 + 3 .. 100

func test01()(sum int){
    for i := 1; i <= 100; i++{
        sum +=i
     }
     return

}

func test02(i int) int{
    if i == 1{
        return 1
    }
    
    return i + test02(i - 1)

} 
func test03(i int) int{
    if i == 100 {
        return 100
    }
    
    return i + test03(i + 1)

}

func main(){
    var sum int
    sum = test01()
    var s int
    s = test02(100)
    
    s1 = test03(1)
    fmt.Println("main")
   

}
```

## 函数类型( 就是所说的 C语言中的 指针）
```markdown
package main

import "fmt"

func Add(a,b int) int{
    return a + b
}

func minus (a ,b int) int {
    return a - b
}

//函数也是一种数据类型，通过type 给一个函数类型起名
// FuncType 它是一个函数类型

type FuncType func(int,int) int  //没有函数名字，没有{}

func main(){
    var result int
    result = Add(1,1)//传统调用方式
    fmt.Println("result = ",result)
    //声明一个函数类型的变量，变量名叫fTest
    var fTest FuncType
    fTest = Add  //是变量就可以赋值
    result = fTest(10,20)  //等价于 Add（10,20)
    fmt.Println("result2  fTest Add = ",result)
    
    fTest = Minus
    result =  fTest(10,5)
    fmt.Println("result3 fTest Minus = ",result)
}
```
## 回调函数
```markdown
package main

import "fmt"
//回调函数，函数有一个参数是函数数类型，这个函数就是回调函数
// 计算器，可以进行四则运算
//多态，多种形态，调用同一个接口，不同的表现，可以实现不同表现，加减乘除
// 现在有想法，后面再实现功能
type FuncType func(int, int) int

//实现加法
func Add(a,b int) int{
    return a + b
}

func Minus(a ,b int) int {
    result a - b
}

func Mul(a,b int) int {
    return a * b

}

func Calc(a,b int, fTest FuncType)  ( result int){
    fmt.Println("Calc")
    result =  fTest(a,b) //这个函数还没有实现
   // result = Add(a,b) //Add()必须先定义后，才能调用
    return
    
}


func main(){

    a := Calc(1,1,Add)
    fmt.Println("a = ",a)
    b := Calc(1,2,Minus)
    fmt.Println("b = ",b)

}

```
## 匿名函数与闭包

```markdown
package main

import "fmt"

func main(){
	a := 10
	str := "mike"
	//匿名函数，没有函数名字，函数定义，还没有调用
	f1 := func(){  // := 自动推导类型
		fmt.Println("a = ",a)
		fmt.Println("str = ",str)
	}
	
	f1()
	// 给一个函数类型起个别名
	type FuncType func()
	
	var f2 FuncType
	f2 = f1
	f2()
	
	// 定义匿名函数，同时调用
	func(){
		fmt.Printf("a = %d, str = %s\n",a,str)
	}() //后面的()代表调用此匿名函数
	
	// 带参数的匿名函数
	f3 := func(i,j int){
		fmt.Prinf("i = %d  j = %d\n",i,j)
	}
	f3(3,4) //只有调用
	
	func(i,j int){
		fmt.Printf("i = %d j = %d\n"i,j)
	}(11,22) //定义匿名函数，同时调用
	
	//匿名函数，有参有返回值
	x,y := func(i,j int)(max,min int){
		if i > j {
			max = i
			min = j
		}else{
			max = j
			min = i
		}
		return
	}(10,20)
	
}

```
## 闭包
```markdown 
pacakge main

import "fmt"

func main(){
	a := 100
	str := "lawrence"
	//匿名函数，没有函数的名字，函数定义，还没有调用
	f1 := func(){ //自动推动类型
		fmt.Println("a = ",a)
		fmt.Println("str = ",str)
	}
	f1() // 调用
	
	//给函数类型起个别名
	type FuncType func()
}

```
