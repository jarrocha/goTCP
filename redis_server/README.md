
## Table of Contents
- [Redis Server](#redis-server)
	- [Redis Server Analysis](#redis-server-analysis)
	- [Redis Server Possible Improvements](#redis-server-possible-improvements)
	- [Redis Server Output](#redis-server-output)

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
