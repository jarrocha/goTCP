package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

type user struct {
	name string
	addr string
	msg  chan string
}

var (
	entering = make(chan user)
	leaving  = make(chan user)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[user]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			fmt.Println(msg)
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.msg <- msg
			}
		case cli := <-entering:
			fmt.Println("Entered:", cli.name)
			clients[cli] = true
		case cli := <-leaving:
			fmt.Println("Leaving:", cli.name)
			delete(clients, cli)
			close(cli.msg)
		}
	}
}

func handleConn(conn net.Conn) {

	// user creation
	var name string
	msg := make(chan string) // outgoing client messages
	addr := conn.RemoteAddr().String()

	// message writer
	go clientWriter(conn, msg)

	// struct setup
	client := user{name, addr, msg}

	// obtain user name
	fmt.Fprintln(conn, "Enter username: ")
	input := bufio.NewScanner(conn)
	for input.Scan() {
		client.name = input.Text()
		break
	}

	// entereed lobby
	messages <- client.name + " has arrived"
	entering <- client

	// scanning for messages
	for input.Scan() {
		messages <- client.name + ":" + input.Text()
	}

	// leaving lobby
	leaving <- client
	messages <- client.name + " has left"
	conn.Close()

}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func main() {
	port := flag.String("p", "8000", "Port number")
	flag.Parse()

	ls, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}
	defer ls.Close()

	go broadcaster()

	for {
		conn, err := ls.Accept()
		if err != nil {
			log.Print(err)
		}

		go handleConn(conn)
	}
}
