package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776 (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424 (exceeds 1 << 64)
	YiB // 1208925819614629174706176
)

func main() {
	//无符号整数
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)

	//
	var i int8 = 127
	fmt.Println(i, i+1, i*i)

	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}

	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i])
	}

	var apples int32 = 1
	var oranges int16 = 2
	//var compote int = apples + oranges
	var compote2 int = int(apples) + int(oranges)
	//fmt.Println(compote)
	fmt.Println(compote2)

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)

	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan)

	var xx complex128 = complex(1, 2)
	var yy complex128 = complex(3, 4)
	fmt.Println(xx * yy)
	fmt.Println(real(xx * yy))
	fmt.Println(imag(xx * yy))
	xxx := 1 + 2i
	yyy := 3 + 4i
	fmt.Println(xxx, yyy)
	fmt.Println(cmplx.Sqrt(-1))

	s := "hello world"
	fmt.Println(len(s))     // "12"
	fmt.Println(s[0], s[7]) // "104 119" ('h' and 'w')
	fmt.Println(s[:5])
	fmt.Println(s[7:])
	fmt.Println(s[:])
	u2 := s[len(s)]
	fmt.Println(u2)

	const GoUsage = `Go is a tool for managing Go source code.
Usage:
go command [arguments]
...`

	sabc := "abc"
	bbb := []byte(sabc)
	sabc2 := string(bbb)
	print(sabc2)

}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func itob(i int) bool {
	return i != 0
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
