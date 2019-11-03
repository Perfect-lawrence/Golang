package main

import (
	"fmt"
)

//golang中的基本数据类型的转换
func main() {
	var t int32 = 100
	// 希望将 t => float
	var t2 float32 = float32(t)
	var t3 int8 = int8(t)
	var t4 int64 = int64(t)
	fmt.Printf("t = %v t2 = %v t3 = %v t4 = %v\n", t, t2, t3, t4)

}

