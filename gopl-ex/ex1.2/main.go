package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {

		fmt.Println("%d:%s\n", i, arg)
	}
}
