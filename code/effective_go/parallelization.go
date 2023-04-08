package main

import "runtime"

type Vector []float64

func (v Vector) Op(i float64) float64 {
	return 1
}

func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		v[i] += u.Op(v[i])
	}
	c <- 1
}

var numCPU2 = runtime.NumCPU()

var numCPU3 = runtime.GOMAXPROCS(0)

const numCPU = 4

func (v Vector) DoAll(u Vector) {
	c := make(chan int, numCPU)
	for i := 0; i < numCPU; i++ {
		go v.DoSome(i*len(v)/numCPU, (i+1)*len(v)/numCPU, u, c)
	}

	for i := 0; i < numCPU; i++ {
		<-c
	}
}

func main() {

}
