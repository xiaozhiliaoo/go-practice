package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func add(x int, y int) int   { return x + y }
func sub(x, y int) (z int)   { z = x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }

func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer2 %d\n", x)
	f(x - 1)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

//func Parse(input string) (s *Syntax, err error) {
//	defer2 func() {
//		if p := recover(); p != nil {
//			err = fmt.Errorf("internal error: %v", p)
//		}
//	}()
//	// ...parser...
//}

func main() {

	defer printStack()
	f(3)

	fmt.Println(triple(4)) // "12"

	_ = double(4)

	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", sub)
	fmt.Printf("%T\n", first)
	fmt.Printf("%T\n", zero)

	f := square
	fmt.Println(f(3))

	f = negative
	fmt.Println(f(3))

	fmt.Printf("%T\n", f) // "func(int) int"

	//f = product

	var ff func(int) int
	//ff(3)
	if ff != nil {
		ff(3)
	}

	fmt.Println(strings.Map(add1, "HAL-9000"))
	fmt.Println(strings.Map(add1, "VMS"))   // "WNT"
	fmt.Println(strings.Map(add1, "Admix")) // "Benjy"

	//匿名函數
	strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000")

}

func add1(r rune) rune { return r + 1 }

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }
