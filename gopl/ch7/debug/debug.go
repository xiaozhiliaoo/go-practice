package main

import (
	"bytes"
	"io"
)

const debug = true

func main() {
	// var buf *bytes.Buffer  //panic
	var buf io.Writer
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
}

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done\n"))
	}
}
