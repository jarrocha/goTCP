## Motivation

This is a collection of small programs written in Golang. These projects are here mostly for inspiration
and quick review of concepts. They will eventually increase in functionality and move to their own repository.

## Table of Contents
- [Redis Server](#redis-server)
	- [Redis Server Analysis](#redis-server-analysis)
	- [Redis Server Possible Improvements](#redis-server-possible-improvements)
	- [Redis Server Output](#redis-server-output)
- [Chat Server](#chat-server)
	- [Chat Server Analysis](#chat-server-analysis)
	- [Chat Server Possible Improvements](#chat-server-possible-improvements)
	- [Chat Server Output](#chat-server-output)
- [TCP Server](#tcp-server)
- [Netcat](#netcat)

## Redis Server
The first version of this server uses a simple approach where each client go about making requests to the database. 
This database is a global variable which leads to a race condition.

To avoid that, the second version uses a single goroutine to handle database access. Another to start the main program 
and one for each connecting client. The approach uses channel for synchronization which is the typical approach using Golang. 
I found it very clever and the uses of channels makes the code very easy to read.

### Redis Server Analysis
The database goroutine resides in the function called `redisServer()` and 
it uses an unbuffered channel for synchronizing actions with multiple clients. So in theory, 
that could also be a performance drawback and could be a good starting point for the third version to improve upon.

### Redis Server Possible Improvements
Use multiple goroutines accessing the database to better scaling and use unbuffered channels to synchronize write 
and delete access to the database between them. No synchronization is needed for read access.
Develop a program to simulate thousands clients with preset amount random commands. Measure the time of response 
for each transaction and obtain a final average.

### Redis Server Output
#### Server output:
```
redis_server2$ go run main.go
Connection from:  127.0.0.1:36118
Performed GET
Performed SET
Performed GET
Performed SET
Performed SET
Performed DEL
Performed GET
Closed Connection from:  127.0.0.1:36118
```

#### Client Output:
```
redis_server2$ netcat localhost 8080
get john
Value not Found
set john 5
New value added
get john
5
set adam 6
New value added
set vick 7
New value added
del john
Value found and deleted
get john
Value not Found
^C
```

## Chat Server
This program presents an interesting approach to using channels and a select statement.

The first version is similar to the one presented in the book the Go Programming Language. 
I saw this program and decided to analyze the thought process behind it deeply since it is a very interesting approach. 
As you can see, the tools to solve problems in this language makes you think from the point of view of scale and performance right from
the start.

There are two versions for this program the first one uses only a string channel for each client, while the second one uses a 
struct with a communications channel and a string for the name for each client. The background server process is the same for both.

### Chat Server Analysis
The core of these programs is a string channel shared between the core chat messaging system and each user. 
This channel is used for each client to send a message to be broad casted and also for the system to send broadcasted 
messages to each user. Each client uses two goroutines, one to wait for input and one to display messages sent by the server.

### Chat Server Possible Improvements
A slow connection from a client can make the server program to get stuck given that it has to wait for a user client to 
read the message before processing others. It is necessary to create a non-blocking mechanism to send messages and probably 
to add a buffer too. This will be added to a third version of the server.

### Chat Server Output
Server output:
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

Client1 output:
```
$ netcat localhost 8000
Enter username:
kyle
john has arrived
john:hello
hello john, you okay?
kyle:hello john, you okay?
john:im good, thanks
john:bye
john has left
bye
kyle:bye
^C
```

Client2 output:
```
$ netcat localhost 8000
Enter username:
john
hello
john:hello
kyle:hello john, you okay?
im good, thanks
john:im good, thanks
bye
john:bye
^C
```

## Netcat











## TCP Server
