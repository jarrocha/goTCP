package main

import (
	"net"
	"flag"
	"log"
	"fmt"
)

type user struct {
	name 	string
	msg		chan message
}

type message struct {
	username string
	text	 string
}

type chatServer struct {
	users 	map[string]user
	join 	chan user
	leave 	chan user
	input 	chan message
}

func (c *chatServer) controller() {
	for {
		select {
		case user := <-c.join:
			c.users[user.name] = user
			fmt.Println(user, " entered the lobby")
		case user := <-c.leave:
			delete(c.users, user.name)
			fmt.Println(user, " left the lobby")
		case msg := <-c.input:
			fmt.Println(msg)
			for _, user := range c.users {
				user.msg<- msg
			}
		}
	}
}

func connHandle(conn net.Conn) {

}

func main() {
	port := flag.String("p", "8000", "Port number")
	flag.Parse()

	li, err := net.Listen("tcp", "localhost:" + *port)
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Print(err)
		}

		go connHandle(conn)

	}
}