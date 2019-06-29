package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	// map是一种无序的基于key-value 的数据结构，Go语言中的map是引用类型，必须初始化才能使用
	// map定义语法： map[keyType]valueType
	// var a map[key] = value
	// 只是声明了一个map类型，但是没有初始化，a就是初始值nil

	var a map[string]int
	fmt.Println(a == nil)

	// map 初始化
	a = make(map[string]int, 8)
	fmt.Println(a == nil)

	//map中添加键值对
	a["lawrence"] = 19
	a["bigdata"] = 200
	fmt.Printf("a: %#v\n", a)
	fmt.Printf("type: %T\n", a)

	// 声明map的同时完成初始化
	b := map[int]bool{
		1: true,
		2: false,
	}
	fmt.Printf("b: %#v\n", b)
	fmt.Printf("type: %T\n", b)

	//	var c map[int]int
	//	c[100] = 200 //c这个map没有初始化，不能直接操作 //error
	//	fmt.Println(c)
	// 判断某个键存不存在
	var stationMap = make(map[string]int, 8)
	stationMap["衡阳到广州"] = 94
	stationMap["永州到深圳"] = 119
	stationMap["株洲到广州"] = 156
	// 判断 永州到广州 在不在 stationMap里
	value, ok := stationMap["永州到广州"]
	fmt.Println(value, ok)
	if ok {
		fmt.Println("永州到广州在stationMap里:", value)
	} else {
		fmt.Println("无此车站")
	}
	fmt.Println("###################")
	// map的遍历
	//map 是无序的，键值对和添加元素顺序无关
	// for range
	for k, v := range stationMap {
		fmt.Printf("map的遍历的key：%s, map的遍历的value: %d\n", k, v)
	}
	// 只遍历map中的key
	for k := range stationMap {
		fmt.Println("只遍历map中的key: ", k)
	}
	//中遍历map中的value
	for _, v := range stationMap {
		fmt.Println("中遍历map中的value: ", v)
	}
	// 删除 株洲到广州 这个键值对
	delete(stationMap, "株洲到广州")
	fmt.Println(stationMap)

	// 按照某个固定顺序遍历map
	var ssMap = make(map[string]int, 100)
	//添加50个键值对
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%2d", i)
		// import "math/rand:
		value := rand.Intn(100) //0~99的随机整数
		ssMap[key] = value
	}
	for k, v := range ssMap {
		fmt.Println(k, v)
	}

	fmt.Println("按照key从小到大的顺序去遍历 ssMap")
	// 1 先取出所有的key存到切片里
	keys := make([]string, 0, 100) // 声明一个切片
	for k := range ssMap {
		keys = append(keys, k)
	}

	// 2 对key做排序
	// import "sort"
	sort.Strings(keys) // keys目前是一个有序的切片

	// 3 按照排序后的key对ssMap
	for _, key := range keys {
		fmt.Println(key, ssMap[key])
	}
	// 复杂类型的map

	// 元素类型为map的切片
	var mapSlice = make([]map[string]int, 8, 8) // 支持完成了切片的初始化
	// [nil nil nil nil nil nil nil]
	fmt.Println(mapSlice[0] == nil)
	// 还需要完成内部map元素的初始化
	mapSlice[0] = make(map[string]int, 8)
	mapSlice[0]["你好吗"] = 100
	fmt.Println(mapSlice)
}

