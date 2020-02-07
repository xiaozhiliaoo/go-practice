package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//var 变量名字 类型 = 表达式
	var s1 string
	fmt.Println(s1)

	var i, j, k int
	var b, f, s = true, 2.3, "four"
	fmt.Println(i, j, k, b, f, s)

	dd := 1
	//var ddd:=1
	fmt.Println(dd)

	i1, j1 := 0, 1
	fmt.Println(i1, j1)

	xx := 1
	p := &xx
	fmt.Println(p, *p)
	*p = 2
	fmt.Println(xx)

	var x, y int
	fmt.Println(&x == &x, &x == &y, &x == nil)

	//var p11 = f()

	fmt.Println(function() == function())

	vv := 1
	incr(&vv)
	fmt.Println(incr(&vv))
	fmt.Println(vv)

	ppp := new(int)
	fmt.Println(ppp, *ppp)
	*ppp = 2
	fmt.Println(*ppp)

}

func function() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}

func test() {
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x)
		}
	}
}

func test2() {
	x := "hello"
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
	}
}

func test3() {
	if x := f(); x == 0 {
		fmt.Println(x)
	} else if y := g(x); x == y {
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}
	//fmt.Println(x, y)  //// compile error: x and y are not visible here
}

func f() int {
	return 1
}

func g(x int) int {
	return 1
}

var cwd string

func init() {
	//cwd , err := os.Getwd()
	//if err != nil {
	//	log.Fatalf("os.Getwd failed: %v", err)
	//}
}

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Printf("Working directory = %s", cwd)
}

func init() {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
}
