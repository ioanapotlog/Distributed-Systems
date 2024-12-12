# Distributed-Systems

## Project 1 - Client-Server Application

This project implements a client-server system in Go, designed to handle multiple clients simultaneously. The clients and server communicate (TCP) by exchanging messages, with the server processing the data using goroutines to handle the requests.

## Features
1. **Client-Server Architecture:**
   - Multiple clients can connect to a single server using TCP sockets.
   - Clients send requests containing his name, the task number and the data array input.
   - The server processes the requests, using goroutines, and then responds to clients.

2. **Configuration File (config.json):**
   - Contains the following parameters:
     - Maximum size of the data array a client can send.
     - Maximum number of concurrent goroutines allowed.
     - Server listening port.

3. **Server Request Handling:**
   - The server includes specific task functions to process client requests.
   - Each function handles a unique operation.
