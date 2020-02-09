package main

import (
	"fmt"
)

func main() {
	//GOMAXPROCS=1 go run main.go
	//GOMAXPROCS=2 go run main.go
	//runtime.GOMAXPROCS(1) //1个操作系统线程
	//runtime.GOMAXPROCS(2) //2个操作系统线程调用go代码
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}
