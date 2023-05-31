package main

import (
	"fmt"
	"os"
)

func main() {
	sum := 100 + 0o10
	fmt.Println(sum)

	file, _ := os.OpenFile("foo", os.O_RDONLY, 0644)
	fmt.Println(file)
}
