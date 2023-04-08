package main

import "bytes"

var freeList = make(chan *bytes.Buffer, 100)

var serverChan = make(chan *bytes.Buffer)

func client() {
	for {
		var b *bytes.Buffer
		select {
		case b = <-freeList:
		default:
			b = new(bytes.Buffer)
		}
		load(b)
		serverChan <- b
	}
}

func load(buffer *bytes.Buffer) {

}

func main() {

}
