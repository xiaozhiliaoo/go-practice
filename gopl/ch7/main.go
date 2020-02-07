package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	var w io.Writer
	//w.Write([]byte("hello"))
	fmt.Printf("%T\n", w) // "<nil>"
	w = os.Stdout
	fmt.Printf("%T\n", w) // "*os.File"
	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w)

	var x1 interface{} = time.Now()
	fmt.Println(x1)
	var x interface{} = []int{1, 2, 3}
	fmt.Println(x == x)
}
