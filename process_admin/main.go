package main

import (
	"fmt"
	"strconv"
	"time"
)

type Process struct {
	Id uint64
	Running bool
	Channel chan string
}

func (proc *Process) Run() {
	i := uint64(0)
	for proc.Running {
		proc.Channel <- "Id: " + strconv.FormatUint(proc.Id, 10) + ": " + strconv.FormatUint(i, 10)
		i = i + 1
		time.Sleep(time.Millisecond * 500)
	}
}

type ProcessAdmin struct {
	Processes []Process
	Print bool
}

func (procAdmin *ProcessAdmin) Add(id uint64) {
	process := Process{Id: id, Running: true, Channel: make(chan string)}
	go process.Run()
	procAdmin.Processes = append(procAdmin.Processes, process)
}

func (procAdmin *ProcessAdmin) Show() {
	for  {
		for _, process := range procAdmin.Processes {
			select {
			case msg := <- process.Channel:
				if procAdmin.Print {
					fmt.Println(msg)
				}
			}
		}
	}
}

func (procAdmin *ProcessAdmin) Remove(id uint64) {
	processes := make([]Process, 0, len(procAdmin.Processes) - 1)
	deleted := false

	for _, process := range procAdmin.Processes {
		if process.Id != id {
			processes = append(processes, process)
		} else {
			process.Running = false
			deleted = true
		}
	}

	if deleted {
		procAdmin.Processes = processes
		fmt.Println("Id", id, "was deleted")
	} else {
		fmt.Println("Id to delete was not found")
	}
}

func main() {
	var option string
	procAdmin := ProcessAdmin{Processes: []Process{}, Print: false}
	var currentId, idToRemove uint64

	for {
		go procAdmin.Show()

		fmt.Print("1. Add process\n2. Show processes\n3. Remove process\n4. Exit\nOption: ")
		fmt.Scanln(&option)
		fmt.Print("\n")

		switch option {
		case "1":
			procAdmin.Add(currentId)
			fmt.Println("Process created with Id:", currentId)
			currentId = currentId + 1
		case "2":
			procAdmin.Print = true
			fmt.Scanln(&option)
			procAdmin.Print = false
		case "3":
			fmt.Scanln(&idToRemove)
			procAdmin.Remove(idToRemove)
		case "4":
			fmt.Println("Finishing gracefully...")
			return
		default:
			fmt.Println("Undefined option")
		}
	}
}
