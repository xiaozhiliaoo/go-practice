package main

import "fmt"

type Request struct {
	args       []int
	f          func([]int) int
	resultChan chan int
}

func sum(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	return
}

func handle1(queue chan *Request) {
	for req := range queue {
		req.resultChan <- req.f(req.args)
	}
}
func main() {
	request := Request{[]int{3, 4, 5}, sum, make(chan int)}
	//clientRequest <- request
	fmt.Printf("answer: %d\n", <-request.resultChan)

}
