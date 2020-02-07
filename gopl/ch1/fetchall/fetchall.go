package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		//start a goroutine
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		//从channel中取出并打印
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s :%v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	// send to channel ch
	ch <- fmt.Sprintf("%.2fs %7 %s", secs, nbytes, url)
}

// go run fetchall.go https://golang.org http://gopl.io https://godoc.org
