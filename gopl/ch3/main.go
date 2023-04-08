package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "李Hello, 世界"
	p := "李"
	suffix := strings.HasPrefix(s, p)
	fmt.Println(suffix)
}
