package main

import "fmt"

// 函数外只能放置标识符(变量、常量、类型)的声明

//程序的入口函数
func main() {
	fmt.Println("Hello World Welcome to Golang")
}

/* 夸平台编译
1、只需要指定目标操作系统的平台和处理器架构即可：

windows 编译出 linux 可执行文件
SET CGO_ENABLED=0 // 禁用CGO
SET GOOS=linux //目标平台是linux
SET GOARCH=amd64 // 目标处理器架构是amd64
然后再执行 go build 命令，得到的是能够在linux平台运行的可执行文件了。

Mac下编译 Linux和Windows平台64 可以执行文件
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=windows GOARCH=amd64  go build

Linux 下编译 Mac 和 Windows平台 64位可执行程序：
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
*/
