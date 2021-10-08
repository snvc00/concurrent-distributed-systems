package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

// Code for process.go

type Process struct {
	Id uint64
	Counter uint64
}

var (
	PROCESS_FROM_SERVER_TO_CLIENT = 1
	PROCESS_FROM_CLIENT_TO_SERVER = 2
)

func server() {
	// Start server listening to port 9999
	server, error := net.Listen("tcp", ":9999")
	if error != nil {
		fmt.Println(error)
		return
	}

	// Init five processes
	processes := make([]Process, 5, 5)
	for i := 0; i < 5; i++ {
		processes[i] = Process{Id: uint64(i), Counter: 0}
	}

	go updateProcesses(&processes)


	// Accept client connections
	for {
		client, error := server.Accept()
		if error != nil {
			fmt.Println(error)
			continue
		}

		go handleClient(client, &processes)
	}
}

func updateProcesses(processes *[]Process) {
	// Update each process in server and log changes
	for {
		for i := 0; i < len(*processes); i++ {
			fmt.Println((*processes)[i].Id, ":", (*processes)[i].Counter)
			(*processes)[i].Counter++
		}
		fmt.Println("-----")
		time.Sleep(time.Millisecond * 500)
	}
}

func handleClient(client net.Conn, processes *[]Process) {
	var connectionOption int

	// Read from client the operation to perform
	error := gob.NewDecoder(client).Decode(&connectionOption)
	if error != nil {
		fmt.Println(error)
		return
	}

	if connectionOption == PROCESS_FROM_SERVER_TO_CLIENT {
		process := &(*processes)[0]
		*processes = (*processes)[1:]

		error := gob.NewEncoder(client).Encode(process)
		if error != nil {
			fmt.Println(error)
			*processes = append(*processes, *process)
			return
		}

		fmt.Println("Log: Process with ID", (*process).Id, "sent to client at", (*process).Counter)

	} else if connectionOption == PROCESS_FROM_CLIENT_TO_SERVER {
		var returnedProcess Process
		error := gob.NewDecoder(client).Decode(&returnedProcess)
		if error != nil {
			fmt.Println(error)
			return
		}

		fmt.Println("Log: Process with ID", returnedProcess.Id, "received at", returnedProcess.Counter)
		*processes = append(*processes, returnedProcess)
	}
}

func main() {
	go server()

	var input string
	fmt.Scanln(&input)
}
