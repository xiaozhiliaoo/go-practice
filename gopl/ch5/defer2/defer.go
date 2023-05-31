package main

import (
	"fmt"
	"os"
	"runtime"
)

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer2 %d\n", x)
	f(x - 1)
}

func main() {
	defer printStack()
	f(3)
}
func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
