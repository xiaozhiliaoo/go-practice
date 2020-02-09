package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		//发送给abort
		abort <- struct{}{}
	}()
	fmt.Println("Commencing countdown.  Press return to abort.")

	select {
	case <-time.After(10 * time.Second):
		//接受abort channel的message
	case <-abort:
		fmt.Println("Lanuch Abort")
		return

	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
