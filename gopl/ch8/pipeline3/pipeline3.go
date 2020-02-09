package main

import "fmt"

//发送
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

//发送和接收
func squarter(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

//接收并打印
func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	//go squarter(naturals, squares)  //导致死锁
	go squarter(squares, naturals)
	printer(squares)

}
