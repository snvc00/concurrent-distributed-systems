package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
)

type File struct {
	Filename string
	Data     []byte
}

type Client struct {
	Id         uint64
	Connection net.Conn
}

type Backup struct {
	Messages  []string
	Filenames []string
}

func onServerShutdown(clients *[]Client, backup *Backup) {
	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-channel

		generateBackup(backup)

		for _, client := range *clients {
			fmt.Println("Sending disconnection signal to client", client.Id)
			gob.NewEncoder(client.Connection).Encode(0)
			client.Connection.Close()
		}

		os.Exit(0)
	}()
}

func generateBackup(backup *Backup) {
	messages := "Messages:\n"
	filenames := "Filenames:\n"
	for _, message := range (*backup).Messages {
		messages += message + "\n"
	}
	for _, filename := range (*backup).Filenames {
		filenames += filename + "\n"
	}
	data := messages + filenames

	err := os.WriteFile("backup_"+strconv.FormatInt(time.Now().Unix(), 10)+".txt", []byte(data), os.ModePerm)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Backup file created")
	}
}

func server() {
	clients := make([]Client, 0)
	var backup Backup
	var currentId uint64

	server, error := net.Listen("tcp", ":9999")
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println("Server running...")
	onServerShutdown(&clients, &backup)

	for {
		connection, error := server.Accept()
		if error != nil {
			fmt.Println(error)
			continue
		}

		client := Client{Id: currentId, Connection: connection}
		clients = append(clients, client)
		currentId += 1

		go handleClient(client, &clients, &backup)
	}
}

func shareReceivedMessage(senderId *uint64, clients *[]Client, message *string) {
	for _, client := range *clients {
		if client.Id != *senderId {
			gob.NewEncoder(client.Connection).Encode(1)
			formated_message := strconv.FormatUint(*senderId, 10) + ": " + *message
			gob.NewEncoder(client.Connection).Encode(&formated_message)
		}
	}
}

func shareReceivedFile(senderId *uint64, clients *[]Client, file *File) {
	for _, client := range *clients {
		if client.Id != *senderId {
			gob.NewEncoder(client.Connection).Encode(1)
			message := strconv.FormatUint(*senderId, 10) + ": " + (*file).Filename
			gob.NewEncoder(client.Connection).Encode(&message)

			clientDir := "client_" + strconv.FormatUint(client.Id, 10)

			if _, err := os.Stat(clientDir); os.IsNotExist(err) {
				err = os.Mkdir(clientDir, os.ModePerm)
				if err != nil {
					fmt.Println(err)
					continue
				}
			}

			path := filepath.Join(clientDir, file.Filename)
			err := os.WriteFile(path, file.Data, os.ModePerm)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func handleClient(client Client, clients *[]Client, backup *Backup) {
	var code int

	for {
		err := gob.NewDecoder(client.Connection).Decode(&code)
		if err != nil {
			fmt.Println(err)
			return
		}

		switch code {
		case 1:
			var message string
			gob.NewDecoder(client.Connection).Decode(&message)
			fmt.Println("Text from client", client.Id, "\b:", message)
			(*backup).Messages = append((*backup).Messages, "Client "+strconv.FormatUint(client.Id, 10)+", "+message)
			shareReceivedMessage(&client.Id, clients, &message)
		case 2:
			var file File
			gob.NewDecoder(client.Connection).Decode(&file)
			fmt.Println("File from client", client.Id, "\b:", file.Filename)
			(*backup).Filenames = append((*backup).Filenames, "Client "+strconv.FormatUint(client.Id, 10)+", "+file.Filename)
			shareReceivedFile(&client.Id, clients, &file)
		case 3:
			fmt.Println("Client", client.Id, "disconnected")

			var position int
			for i, c := range *clients {
				if client.Id == c.Id {
					position = i
					break
				}
			}

			(*clients)[position] = (*clients)[0]
			*clients = (*clients)[1:]
		}
	}
}

func main() {
	go server()
	
	for {
		fmt.Print("")
	}
}
