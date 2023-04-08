package main

import "fmt"

func main() {
	var cwd string
	if true {
		cwd := "11"
		fmt.Println("1:", cwd)
		cwd = "222"
		fmt.Println("2:", cwd)
	}

	fmt.Println("3:", cwd)
}
