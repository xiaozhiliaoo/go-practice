package main

import "fmt"

func sequares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	fmt.Println(sequares())
	fmt.Println(sequares())
	fmt.Println(sequares())
	fmt.Println(sequares())
	f := sequares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
