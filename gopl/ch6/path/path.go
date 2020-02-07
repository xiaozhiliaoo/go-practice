package main

import (
	"fmt"
	"go-practice/gopl/ch6/geometry"
)

type Path []geometry.Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		println(i)
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func main() {
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())
}
