package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	_, err := os.Open("/no/such/file")
	fmt.Println(err) // "open /no/such/file: No such file or directory"
	fmt.Printf("%#v\n", err)

	fmt.Println(os.IsNotExist(err))

	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)
	c := w.(*bytes.Buffer)
	fmt.Println(f)
	fmt.Println(c)

	//var w io.Writer
	//w = os.Stdout
	//w = new(bytes.Buffer)
	//w = nil
	//
	//var w io.Writer
	//w = os.Stdout
	//w = new(bytes.Buffer)
	//w = time.Second
	//
	//var rwc io.ReadWriteCloser
	//rwc = os.Stdout
	//rwc = new(bytes.Buffer)
}

func sqlQuote(x interface{}) string {
	if x == nil {
		return "NULL"
	} else if _, ok := x.(int); ok {
		return fmt.Sprintf("%d", x)
	} else if _, ok := x.(uint); ok {
		return fmt.Sprintf("%d", x)
	} else if b, ok := x.(bool); ok {
		if b {
			return "TRUE"
		}
		return "FALSE"
	} else if s, ok := x.(string); ok {
		return sqlQuoteString(s) // (not shown)
	} else {
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}

func sqlQuote2(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x) // x has type interface{} here.
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return sqlQuoteString(x) // (not shown)
	default:
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}
