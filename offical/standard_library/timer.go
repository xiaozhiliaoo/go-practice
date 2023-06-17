package main

import (
	"fmt"
	"strconv"
	"time"
)

func task1() {
	ticker := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-ticker.C:
			i += 1
			fmt.Println("Task executed at:", time.Now().String()+"---"+strconv.Itoa(i))
		}
	}
}

func task2() {
	ticker := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-ticker.C:
			i += 1
			fmt.Println("Task2 executed at:", time.Now().String()+"---"+strconv.Itoa(i))
		}
	}
}

var i = 1

func main() {
	go task1()
	go task2()
	time.Sleep(5 * time.Minute)
}
