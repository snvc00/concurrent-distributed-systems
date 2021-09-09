package main

import (
	. "files_sort/process"
	"fmt"
	"os"
	"sort"
)

func main() {
	stringSorting()
	processSorting()
}

func stringSorting() {
	// Size for string slice from standard input
	var size int
	fmt.Print("String slice size: ")
	fmt.Scanln(&size)

	// Create the slice with the given size
	strings := make([]string, size)
	var tmp string
	for i := 0; i < size; i++ {
		fmt.Scanln(&tmp)
		strings[i] = tmp
	}
	fmt.Println(strings)

	// Ascending sort to ascending.txt
	sort.Strings(strings)
	file, err := os.Create("ascending.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, str := range strings {
		file.WriteString(str + "\n")
	}
	file.Close()

	// Descending sort to descending.txt
	sort.Sort(sort.Reverse(sort.StringSlice(strings)))
	file, err = os.Create("descending.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, str := range strings {
		file.WriteString(str + "\n")
	}
	file.Close()
}

func processSorting() {
	// Size for process slice from standard input
	var size int
	fmt.Print("Process slice size: ")
	fmt.Scanln(&size)

	// Create the slice with the given size
	processes := make([]Process, size)
	var id, time uint64
	var priority int64
	var status string
	for i := 0; i < size; i++ {
		fmt.Scanln(&id, &priority, &time, &status)
		processes[i] = Process{Id: id, Priority: priority, Time: time, Status: status}
	}

	// Ascending sort to standard output
	sort.Sort(ByPriority(processes))
	fmt.Println("Ascending by PRIORITY")
	for _, process := range processes {
		fmt.Println(process.String())
	}

	// Descending sort to standard output
	sort.Sort(sort.Reverse(ByPriority(processes)))
	fmt.Println("\nDescending by PRIORITY")
	for _, process := range processes {
		fmt.Println(process.String())
	}
}
