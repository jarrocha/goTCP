## Motivation

This is a collection of small programs written in Golang. These projects are here mostly for inspiration
and quick review of concepts. They will eventually increase in functionality and move to their own repository.

## Table of Contents
- [Chat Server](#chat-server)
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















## Netcat












## Redis Server












## TCP Server
