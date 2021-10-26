package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strconv"
)

type Teacher struct {
	Subjects map[string]map[string]float64
	Students map[string]map[string]float64
}

func PrintSubjects(teacher *Teacher) {
	subjectsJson, err := json.MarshalIndent((*teacher).Subjects, "", "    ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("---\n\"subjects\":", string(subjectsJson), "\n---")
}

func (teacher *Teacher) RegisterStudentGrade(args *map[string]string, reply *string) error {
	studentName := (*args)["name"]
	subjectName := (*args)["subject"]
	grade, err := strconv.ParseFloat((*args)["grade"], 32)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if (*teacher).Subjects[subjectName] == nil {
		(*teacher).Subjects[subjectName] = make(map[string]float64)
	}

	if (*teacher).Students[studentName] == nil {
		(*teacher).Students[studentName] = make(map[string]float64)
	}

	((*teacher).Subjects[subjectName])[studentName] = grade
	((*teacher).Students[studentName])[subjectName] = grade

	PrintSubjects(teacher)

	*reply = "OK"
	return nil
}

func (teacher *Teacher) GetStudentGPA(name *string, reply *float64) error {
	studentSubjects := (*teacher).Students[*name]
	var grades float64

	if studentSubjects == nil {
		return errors.New("No student with name " + *name)
	}

	for _, grade := range studentSubjects {
		grades += grade
	}

	*reply = grades / float64(len(studentSubjects))
	return nil
}

func (teacher *Teacher) GetStudentsAverageGPA(args *int, reply *float64) error {
	var gpaSum float64

	if len((*teacher).Students) == 0 {
		return errors.New("No students registered")
	}

	for _, studentSubjects := range (*teacher).Students {
		grades := 0.0
		for _, grade := range studentSubjects {
			grades += grade
		}
		gpaSum += grades / float64(len(studentSubjects))
	}

	*reply = gpaSum / float64(len((*teacher).Students))
	return nil
}

func (teacher *Teacher) GetSubjectGradesAverage(subjectName *string, reply *float64) error {
	subjectStudents := (*teacher).Subjects[*subjectName]
	var grades float64

	if subjectStudents == nil {
		return errors.New("No subject with name " + *subjectName)
	}

	for _, grade := range subjectStudents {
		grades += grade
	}

	*reply = grades / float64(len(subjectStudents))
	return nil
}

func main() {
	fmt.Println("Server starting...")
	teacher := Teacher{Subjects: make(map[string]map[string]float64), Students: make(map[string]map[string]float64)}

	rpc.Register(&teacher)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal("Listen error:", err)
	}
	go http.Serve(listener, nil)

	fmt.Println("Server running...")
	var input string
	fmt.Scanln(&input)
}
