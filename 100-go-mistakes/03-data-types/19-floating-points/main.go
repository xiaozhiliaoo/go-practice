package main

import (
	"fmt"
	"math"
)

func main() {
	var n float32 = 1.0001
	fmt.Println(n * n)

	fmt.Println(math.SmallestNonzeroFloat32)
	fmt.Println(math.SmallestNonzeroFloat64)

	var a float64
	fmt.Println("Default float:", a)
	positiveInf := 1 / a
	negativeInf := -1 / a
	nan := a / a
	fmt.Println(positiveInf, negativeInf, nan)
	fmt.Println(math.IsNaN(nan))
	fmt.Println(math.IsInf(nan, 1))

	var a1 float32 = 1.0001
	var a2 float32 = 1.0001
	fmt.Println(a1 == a2)

	var b1 float32 = 1.00000001
	var b2 float32 = 1.00000001000001
	fmt.Println(b1 == b2)

	var c1 float32 = 1.001
	var c2 float32 = 1.001001
	fmt.Println(c1 == c2)

	//先加常数还是后加常数
	fmt.Println(f1(100000))
	fmt.Println(f2(100000))

}

func f1(n int) float64 {
	result := 10_000.
	for i := 0; i < n; i++ {
		result += 1.0001
	}
	return result
}

func f2(n int) float64 {
	result := 0.
	for i := 0; i < n; i++ {
		result += 1.0001
	}
	return result + 10_000.
}
