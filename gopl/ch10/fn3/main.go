package main

import (
	//导入包重命名
	fn1 "go-practice/gopl/ch10/fn1"
	fn2 "go-practice/gopl/ch10/fn2"
)

func main() {
	//路径不一样，但是包名一样，路径并不是包名,包名一般是最后一段
	fn1.Fn1()
	fn2.Fn2()
}
