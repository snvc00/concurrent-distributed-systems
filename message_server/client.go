package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type File struct {
	Filename string
	Data     []byte
}

func onClientShutdown(server *net.Conn) {
	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-channel

		gob.NewEncoder(*server).Encode(3)
		(*server).Close()

		os.Exit(0)
	}()
}

func serverListener(server *net.Conn) {
	var code int
	var data string

	for {
		err := gob.NewDecoder(*server).Decode(&code)
		if err != nil {
			fmt.Println(err)
			return
		}

		switch code {
		case 0:
			fmt.Println("Server sent disconnection signal")
			os.Exit(0)
		case 1:
			err = gob.NewDecoder(*server).Decode(&data)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("\n-> Message from client", data, "<-")
		default:
			fmt.Println("Not recognized code")
		}

	}
}

func main() {
	// Open connection with server
	fmt.Println("Connecting to server...")
	server, error := net.Dial("tcp", ":9999")
	if error != nil {
		fmt.Println(error)
		return
	}

	onClientShutdown(&server)
	go serverListener(&server)

	fmt.Println("Connected to server")
	var filename, option string

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\nWelcome!\n1. Send message\n2. Send file\nOption: ")
		fmt.Scanln(&option)

		switch option {
		case "1":
			fmt.Print("Message: ")
			scanner.Scan()
			message := scanner.Text()

			gob.NewEncoder(server).Encode(1)
			gob.NewEncoder(server).Encode(&message)
		case "2":
			fmt.Print("Filename: ")
			fmt.Scanln(&filename)

			data, err := os.ReadFile(filename)
			if err != nil {
				fmt.Println(err)
				continue
			}

			gob.NewEncoder(server).Encode(2)
			gob.NewEncoder(server).Encode(File{Filename: filename, Data: data})
		default:
			fmt.Println("Invalid option")
		}
	}
}
