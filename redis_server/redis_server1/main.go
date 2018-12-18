// This program works concurrently but suffers from race condition
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
)

// database struct
var db = make(map[string]string, 20)

// channel semaphore
var setSem = make(chan struct{})
var delsem = make(chan struct{})

func invalidCommand(conn net.Conn) {
	fmt.Fprintln(conn, "Invalid Command")
}

func processCmd(conn net.Conn, cmd []string) {
	if !(len(cmd) > 0) {
		return
	}

	comm := strings.ToLower(cmd[0])

	switch {
	case comm == "get":
		// GET <KEY>
		fmt.Println("Performed GET")
		if v, found := db[cmd[1]]; !found {
			fmt.Fprintln(conn, "Value not found")
		} else {
			fmt.Fprintln(conn, v)
		}
	case comm == "set":
		// SET <KEY> <VALUE>
		if len(cmd) < 3 {
			invalidCommand(conn)
			return
		}
		fmt.Println("Performed SET")
		setSem <- struct{}{}
		db[cmd[1]] = cmd[2]
		<-setSem
	case comm == "del":
		// DEL <KEY>
		fmt.Println("Performed DEL")
		delete(db, cmd[1])
	case comm == "print":
		fmt.Println("Performed Print")
		fmt.Fprintln(conn, db)
	default:
		invalidCommand(conn)
	}

	return

}

func connHandle(conn net.Conn) {

	defer exit(conn)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		rawcmds := scanner.Text()
		cmds := strings.Fields(rawcmds)

		if cmds[0] == "q" || cmds[0] == "Q" {
			exit(conn)
			return
		}

		if len(cmds) < 2 {
			invalidCommand(conn)
			continue
		} else {
			processCmd(conn, cmds)
		}
	}
}

func exit(conn net.Conn) {
	fmt.Println("Closed Connection from: ", conn.RemoteAddr())
	fmt.Fprintln(conn, "Exiting")
	conn.Close()
}

func main() {

	port := flag.String("p", "8080", "Port number")
	flag.Parse()

	ls, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ls.Accept()
		if err != nil {
			log.Print(err)
		}

		fmt.Println("Connection from: ", conn.RemoteAddr())
		go connHandle(conn)
	}
}
