package main

import "fmt"

// Go九九乘法表

func main() {
	var key = 0
	for i := 1; i < 10; i++ {
		key++
		for j := key; j < 10; j++ {
			if j != 9 {
				fmt.Printf("%d * %d = %d  ", i, j, i*j)
			} else if j == 9 {
				fmt.Printf("%d * %d = %d\n", i, j, i*j)
			}
		}
	}

	// go语言打印九九乘法表
	var t int = 10
	for i := 1; i < t; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %2d\t", i, j, i*j)
		}
		fmt.Println()
	}
}
