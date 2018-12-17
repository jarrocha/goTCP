// Simple TCP netcat program

package main

import (
	"net"
	"os"
	"fmt"
	"log"
	"io"
)

func handleConn(conn net.Conn) {
	fmt.Println("Connected to: ", conn.RemoteAddr())
	go io.Copy(os.Stdout, conn)
	if n, _ := io.Copy(conn, os.Stdin); n > 0 {
		fmt.Println(n, "bytes copied")
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./netcat ADDR PORT")
		return
	}

	addr := os.Args[1]
	port := os.Args[2]

	conn, err := net.Dial("tcp", addr + ":" + port)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	handleConn(conn)
}