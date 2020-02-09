package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	//main gorouites
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	fmt.Println("handleConn:", conn)
	//双向channel
	ch := make(chan string)
	go clientWriter(conn, ch)
	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	//客户端输入的字符
	input := bufio.NewScanner(conn)
	for input.Scan() {
		fmt.Println("input scan:")
		messages <- who + ": " + input.Text()
	}

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

//ch是接收channel
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		//返回给客户端
		fmt.Println("clientWriter ...")
		fmt.Fprintln(conn, "~~~", msg)
	}
}

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	//全局消息
	messages = make(chan string)
)

func broadcaster() {
	fmt.Println("broadcaster:")
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}
