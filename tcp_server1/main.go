// Example of a hello world TCP server
// Use netcat or telnet on localhost:8080 to access

package main

import (
	"net"
	"log"
	"fmt"
	"time"
)

func handleConn(conn net.Conn) {
	fmt.Println("Connection from: ", conn.RemoteAddr())

	for x := 0; x < 3 ; x++ {
		fmt.Fprintf(conn, "Hello World\n")
		time.Sleep(1 * time.Second)
	}

	fmt.Fprintf(conn, "Bye\n")
	conn.Close()
	
}

func main() {

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print(err)
		}

		handleConn(conn)
		fmt.Println("Connection Closed")
	}

}