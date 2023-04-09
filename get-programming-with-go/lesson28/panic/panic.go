package main

import "fmt"

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("catch:", e)
		}
	}()

	panic("I forgot my towel")
}
