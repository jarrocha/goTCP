// Example of an echo server
// Use netcat or telnet on localhost:8080 to access

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"time"
	"strings"
)

func echo(c net.Conn, msg string, delay time.Duration) {
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToUpper(msg))
}

func handleConn(conn net.Conn) {
	remaddr := conn.RemoteAddr()
	fmt.Println("Connection from: ", remaddr)

	msg := bufio.NewScanner(conn)
	for msg.Scan() {
		go echo(conn, msg.Text(), 1 * time.Second)
	}
	conn.Close()
	fmt.Println("Connection from: ", remaddr, " closed.")
}

func main() {
	port := flag.String("p", "8080", "Port number")
	flag.Parse()

	server := "localhost:" + *port

	ls, err := net.Listen("tcp", server)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ls.Accept()
		if err != nil {
			log.Print(err)
		}

		go handleConn(conn)
	}
}
