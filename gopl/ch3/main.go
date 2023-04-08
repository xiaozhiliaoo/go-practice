package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
		i += size
	}

	fmt.Println()

	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	s1 := "Hello, 世界"
	fmt.Println(s1[1:9])

}
