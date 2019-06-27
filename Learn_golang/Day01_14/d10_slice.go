package main

import (
	"fmt"
	"sort"
)

func main() { // 切片是个引用类型,切片不能直接比较，可以通过len()长度与 0 比较
	//切片（slice),其实底层就是数组，切片不要定义 数组的长度
	//声明切片类型
	var a []string              //声明一个字符串的切片
	var b []int                 //声明一个整形切片并初始化
	var c = []bool{false, true} //声明一个布尔类型切片并初始化
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	//基于数组的切片
	aa := [5]int{1, 2, 3, 5, 6}
	bb := aa[1:4]
	fmt.Println(bb)
	fmt.Printf("%T\n", bb)
	//切片再次切片
	cc := bb[0:len(bb)]
	fmt.Println(c)
	fmt.Printf("cc: %T\n", cc)
	// make 函数构造切片
	// make([]类型,长度（元素个数）,容量(可以省略))
	// make([]data_type,size(len),cap)
	dd := make([]int, 5, 10)
	fmt.Println(dd)
	fmt.Printf("dd: %T\n", dd)
	//通过len()函数获取切片的长度
	fmt.Println(len(dd))
	//通过cap()函数获取切片的容量
	fmt.Println(cap(dd))

	// nil 是一个引用类型
	var ff []int //声明一个int类型切片
	if ff == nil {
		fmt.Println("a==nil")
	}
	fmt.Println(ff, len(ff), cap(ff))

	var ee = []int{} //声明并且初始化
	if ee == nil {
		fmt.Println("ee==nil")
	}
	fmt.Println(ee, len(ee), cap(ee))

	gg := make([]int, 0)
	if gg == nil {
		fmt.Println("gg==nil")
	}
	fmt.Println(gg, len(gg), cap(gg))

	hh := make([]int, 3) //[0 0 0]
	uu := hh             // 其实uu 和 hh 是共用同一个底层数组
	uu[0] = 100
	fmt.Println(hh) // [0 0 0] 改变了[100 0 0]
	fmt.Println(uu)
	// 切片的遍历 支持索引 遍历 和 range 遍历
	//根据索引来遍历
	oo := []int{1, 3, 5, 6, 7}
	for i := 0; i < len(oo); i++ {
		fmt.Println(oo, oo[i])
	}
	fmt.Println("###################")
	// for range 遍历
	for index, value := range oo {
		fmt.Println(index, value)
	}
	//append()方法为切片添加元素
	//切片的扩容
	//切片要初始化之后才能使用
	var pp []int //此时它并没又申请内存
	pp = append(pp, 10)
	fmt.Println(pp)
	for i := 0; i < 10; i++ {
		pp = append(pp, i)
		fmt.Printf("%v len:%d cap:%d ptr:%p\n", pp, len(pp), cap(pp), pp)
	}
	// 一次追加多个元素
	var yy []int
	ww := []int{11, 22, 33, 44, 55}
	yy = append(yy, 1, 3, 5, 6, 9, 12)
	fmt.Println(yy)
	ww = append(ww, yy...) //把yy里面的元素也添加进来
	fmt.Println(ww)

	//copy()函数复制切片
	// copy(srtslice,destslice)
	rr := []int{11, 33, 44, 55, 77}
	xx := make([]int, 5, 6)
	vv := xx // vv := xx 这里两个都是指向同一个底层数组
	copy(xx, rr)
	vv[0] = 200
	fmt.Println("rr: ", rr)
	fmt.Println("xx: ", xx)
	fmt.Println("vv: ", vv)
	// 切片删除元素
	qq := []string{"北京", "上海", "广州", "深圳", "成都"}
	qq = append(qq[0:2], qq[3:]...)
	fmt.Println(qq)
	// append(qq[:index],qq[index+1:]...)

	// 从切片中删除元素
	nn := []int{40, 50, 70, 80, 90}
	//要删除索引为2的元素
	nn = append(nn[:2], aa[3:]...)
	fmt.Println("nn = ", nn)

	var mm = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		mm = append(mm, fmt.Sprintf("%v", i))
	}
	fmt.Println("mm: ", mm) //[      0 1 2 3 4 5 6 7 8 9]
	tt := []int{4, 5, 6, 1, 9, 7, 8, 2}
	// import sort
	sort.Ints(tt[:])
	fmt.Println("Sort.Ints(tt[:]) = ", tt)
}
