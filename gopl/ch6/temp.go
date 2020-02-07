package main

import (
	"fmt"
	"go-practice/gopl/ch6/geometry"
)

func main() {
	p := geometry.Point{1, 2}
	q := geometry.Point{4, 6}
	fmt.Println(geometry.Distance(p, q))
	fmt.Println(p.Distance(q))

	r := &geometry.Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)

	p1 := geometry.Point{1, 2}
	pptr := &p1
	pptr.ScaleBy(2)
	fmt.Println(p1)

	p2 := geometry.Point{1, 2}
	(&p2).ScaleBy(2)
	fmt.Println(p2)

	p3 := geometry.Point{1, 2}
	p3.ScaleBy(2)
	fmt.Println(p3)

	p4 := geometry.Point{1, 2}
	p4.ScaleBy2(2)
	fmt.Println(p4)

	p5 := geometry.Point{1, 2}
	(&p5).ScaleBy2(2)
	fmt.Println(p5)

}
