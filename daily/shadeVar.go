package main

import "fmt"

func main() {
	var str string
	if true {
		str := "111"
		if true {
			str = "111"
			fmt.Println("true1:" + str)
		}
		fmt.Println("true2:" + str)
	}
	fmt.Println("false:" + str)
}
