package main

import "net/http"

const MaxOutStanding = 100

var sem = make(chan int, MaxOutStanding)

func handle(r *http.Request) {
	sem <- 1
	process(r)
	<-sem
}

func process(r *http.Request) {

}

func Serve(queue chan *http.Request) {
	for {
		req := <-queue
		go handle(req)
	}
}

func Serve2(queue chan *http.Request) {
	for req := range queue {
		sem <- 1
		go func() {
			process(req)
			<-sem
		}()
	}
}

func Serve3(queue chan *http.Request) {
	for req := range queue {
		sem <- 1
		go func(req *http.Request) {
			process(req)
			<-sem
		}(req)
	}
}

func Serve4(queue chan *http.Request) {
	for req := range queue {
		req := req // Create new instance of req for the goroutine.
		sem <- 1
		go func() {
			process(req)
			<-sem
		}()
	}
}

func main() {

}
