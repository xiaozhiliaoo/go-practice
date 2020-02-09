package main

import (
	"fmt"
	"time"
)

func main() {

	tick := time.Tick(1 * time.Second)
	<-tick

	ch := make(chan string, 3)
	ch <- "A"
	ch <- "B"
	ch <- "C"
	fmt.Println(len(ch))
	fmt.Println(cap(ch))
	fmt.Println(<-ch)
	fmt.Println(len(ch))
	fmt.Println(<-ch)
	fmt.Println(len(ch))
	fmt.Println(<-ch)
	fmt.Println(cap(ch))

	fmt.Println("---------error occur --------")
	fmt.Println(<-ch)

	//fatal error: all goroutines are asleep - deadlock!
	//ch2:=make(chan string)
	//ch2 <- "A"
	//ch2 <- "B"
	//ch2 <- "C"
	//fmt.Println(<-ch2)
}
