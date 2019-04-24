# Linux安装Go语言环境

* 安装 依赖包

```
yum install mercurial git gcc-c++ gcc openssl readline readline-devel zlib zlib-deve openssl openssl-devel wget -y
wget https://studygolang.com/dl/golang/go1.12.4.linux-amd64.tar.gz
```
* 下载完成 用tar 命令来解压压缩包

```
tar -zxvf go1.12.4.linux-amd64.tar.gz -C /usr/local/
```
* 建立Go语言的工作空间（workspace，也就是GOPATH环境变量指向的目录）
Go语言代码必须在工作空间内。工作空间是一个目录，其中包含三个子目录：
src ---- 里面每一个子目录，就是一个包。包内是Go语言的源码文件
pkg ---- 编译后生成的，包的目标文件
bin ---- 生成的可执行文件
这里，我们在/home目录下，建立一个名为go（可以不是go, 任意名字都可以）的文件夹，然后再建立三个子文件夹（子文件夹名必须为src、pkg、bin）。

```
echo "GOROOT=/usr/local/go" >> /etc/profile
echo "GOPATH=/home/lawrence/Go_Project/calcproj" >> /etc/profile
echo "export PATH=$GOROOT/bin:$GOPATH:$PATH" >> /etc/profile
source /etc/profile
mkdir /home/lawrence/Go_Project/mysql/{bin,pkg,src}
[lawrence@izwz99f5wnpdp]$ go version
go version go1.12.4 linux/amd64
[lawrence@izwz99f5wnpdp]$ go env
```

# go hello world

```
cd /home/lawrence/Go_Project/mysql/src

cat first.go

package main

import "fmt"

func main() {

	fmt.Println("hello Golang")

}

```
```
[lawrence@izwz99f5wnpdp]$ go fmt first.go # 格式化代码格式
[lawrence@izwz99f5wnpdp]$ go run first.go
hello Golang
```
