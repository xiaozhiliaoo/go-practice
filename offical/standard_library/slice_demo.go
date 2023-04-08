package main

import (
	"fmt"
)

// GOOS=linux GOARCH=amd64 go tool compile -S -N -l slice_demo.go > slice_demo.s
func main() {
	a := make([]int, 0, 2)
	for i := 0; i < 20; i++ {
		a = append(a, i)
		fmt.Println(len(a), cap(a), a)
	}
}
