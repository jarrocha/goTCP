## Motivation

This is a collection of small programs written in Golang. These projects are here mostly for inspiration
and quick review of concepts. They will eventually increase in functionality and move to their own repository.

## Table of Contents
- [Chat Server](#chat-server)
	- [Brief Analysis](#brief-analysis)
	- [Possible Improvements](#possible-improvements)
- [Netcat](#netcat)
- [Redis Server](#redis-server)
- [TCP Server](#tcp-server)



## Chat Server
This program presents an interesting approach to using channels and a select statement.

The first version is similar to the one presented in the book the Go Programming Language. 
I saw this program and decided to analyze the thought process behind it deeply since it is a very interesting approach. 
As you can see, the tools to solve problems in this language makes you think from the point of view of scale and performance right from
the start.

There are two versions for this program the first one uses only a string channel for each client, while the second one uses a 
struct with a communications channel and a string for the name for each client. The background server process is the same for both.

### Brief Analysis
The core of these programs is a string channel shared between the core chat messaging system and each user. 
This channel is used for each client to send a message to be broad casted and also for the system to send broadcasted 
messages to each user. Each client uses two goroutines, one to wait for input and one to display messages sent by the server.

### Possible Improvements
A slow connection from a client can make the server program to get stuck given that it has to wait for a user client to 
read the message before processing others. It is necessary to create a non-blocking mechanism to send messages and probably 
to add a buffer too. This will be added to a third version of the server.

### Output
Server Output:
```
chat_server2$ go run main.go
kyle has arrived
Entered: kyle
john has arrived
Entered: john
john:hello
kyle:hello john, you okay?
john:im good, thanks
john:bye
Leaving: john
john has left
kyle:bye
Leaving: kyle
kyle has left
```


## Netcat












## Redis Server












## TCP Server
