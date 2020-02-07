package main

import (
	"log"
	"time"
)

func trace(msg string) func() {
	start := time.Now()
	log.Print("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	time.Sleep(1 * time.Second)
}

func main() {
	bigSlowOperation()
}
