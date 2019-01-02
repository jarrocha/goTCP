// Version2 that eliminates race conditions

/* 	Description:
    Though still a very simplistic approach to the problem. It is a step forward in how to divide a problem
	into smaller pieces (go routines) and dispatch requests. This current approach is a one to many processing
	of request. There are two channels used for communication. The first one is used to let the redisServer()
	function know there is a new command input, the second one is to let the user go routine know there's a result
	to the request. No matter how many user request are done, there's only one go routine serving all of them at
	a time. Sort of a desorganized queue. There's no guarantee that if your request arrives second, it will be
	served as second.
	Improvement:
	One possible improvement is to have many to many approach and use a channel to avoid race condition on the
	go routines performing database access. Specifically for the delete and set commands.
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

// response strings
var valNotFound = "Value not Found\n"
var valInva = "Invalid command\n"
var valDel = "Value found and deleted\n"
var valUpd = "Value updated\n"
var valNew = "New value added\n"
var valnoSpace = "\n"

// Command struct
type Command struct {
	Fields []string
	Result chan string
}

func redisServer(cmds chan Command) {
	var db = make(map[string]string, 20)
	for cmd := range cmds {
		if len(cmd.Fields) < 2 {
			cmd.Result <- valInva
			return
		}

		switch cmd.Fields[0] {
		case "get":
			// GET <KEY>
			fmt.Println("Performed GET")
			key := cmd.Fields[1]
			val, found := db[key]
			if !found {
				cmd.Result <- valNotFound
				break
			}
			cmd.Result <- val + "\n"
		case "set":
			// SET <KEY> <VALUE>
			if len(cmd.Fields) < 3 {
				cmd.Result <- valInva
				return
			}

			fmt.Println("Performed SET")
			key := cmd.Fields[1]
			val := cmd.Fields[2]
			_, found := db[key]
			db[key] = val
			if found {
				cmd.Result <- valUpd
			} else {
				cmd.Result <- valNew
			}
		case "del":
			// DEL <KEY>
			fmt.Println("Performed DEL")
			key := cmd.Fields[1]
			_, found := db[key]
			if found {
				delete(db, key)
				cmd.Result <- valDel
			} else {
				cmd.Result <- valNotFound
			}
		default:
			cmd.Result <- valInva
		}
	}
}

func invalidCommand(conn net.Conn) {
	fmt.Fprintln(conn, "Invalid Command")
}

func connHandle(commands chan Command, conn net.Conn) {

	defer exit(conn)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		rawcmds := scanner.Text()
		cmds := strings.Fields(rawcmds)

		result := make(chan string)
		commands <- Command{
			Fields: cmds,
			Result: result,
		}

		io.WriteString(conn, <-result)
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
	defer ls.Close()

	commands := make(chan Command)
	go redisServer(commands)

	for {
		conn, err := ls.Accept()
		if err != nil {
			log.Print(err)
		}

		fmt.Println("Connection from: ", conn.RemoteAddr())
		go connHandle(commands, conn)
	}
}
