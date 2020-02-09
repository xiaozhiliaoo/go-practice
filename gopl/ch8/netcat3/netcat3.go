package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	//background goroutine
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		//发送struct{}{}数据给done channel
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	//从main goroutine 中接收done channel中的数据，不接收值会阻塞住
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
