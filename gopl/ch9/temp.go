package main

import (
	"errors"
	"fmt"
)

func main() {
	//var x []int
	//go func() { x = make([]int, 10) }()
	//go func() { x = make([]int, 1000000) }()
	//x[999999] = 1

	errors.New()

	var x, y int
	go func() {
		x = 1                   // A1
		fmt.Print("y:", y, " ") // A2
	}()
	go func() {
		y = 1                   // B1
		fmt.Print("x:", x, " ") // B2
	}()
}
