package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Code for process.go

type Process struct {
	Id      uint64
	Counter uint64
}

var (
	PROCESS_FROM_SERVER_TO_CLIENT = 1
	PROCESS_FROM_CLIENT_TO_SERVER = 2
)

// Listen for interrupt (Ctrl + C) and send process back to server
func onClientShutdown(process *Process) {
	// Create a channel to listen for OS interruption
	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)

	go func() {
		<- channel

		fmt.Println("Log: Sending process back to server with ID", (*process).Id)

		connection, error := net.Dial("tcp", ":9999")
		if error != nil {
			fmt.Println(error)
			return
		}

		// Send request to return the process and send it
		gob.NewEncoder(connection).Encode(PROCESS_FROM_CLIENT_TO_SERVER)
		gob.NewEncoder(connection).Encode(process)
		connection.Close()

		// Exit successfully
		os.Exit(0)
	}()
}

func main() {
	// Open connection with server
	connection, error := net.Dial("tcp", ":9999")
	if error != nil {
		fmt.Println(error)
		return
	}

	// Send message to request a process
	error = gob.NewEncoder(connection).Encode(PROCESS_FROM_SERVER_TO_CLIENT)
	if error != nil {
		fmt.Println(error)
	}

	// Receive process
	var process Process
	error = gob.NewDecoder(connection).Decode(&process)
	if error != nil {
		fmt.Println(error)
		return
	}

	// Prepare for client shutdown
	onClientShutdown(&process)
	connection.Close()

	// Update process data and log changes
	for {
		fmt.Println(process.Id, ":", process.Counter)
		process.Counter++
		time.Sleep(time.Millisecond * 500)
	}
}
