package main

import (
	"bytes"
	"fmt"
	"sync"
)

type SyncBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

func main() {
	p := new(SyncBuffer)
	fmt.Printf("%T", p)

}
