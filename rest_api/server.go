package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Student struct {
	Id       uint64             `json:"id"`
	Name     string             `json:"name"`
	Subjects map[string]float64 `json:"subjects"`
}

var studentsGroup map[uint64]Student
var nextId int

func students(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method, req.URL.Path)
	res.Header().Set("Content-Type", "application/json")

	switch req.Method {
	case "GET":
		studentsJson, err := json.MarshalIndent(studentsGroup, "", "    ")
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(res, string(studentsJson))
	case "POST":
		var student Student
		err := json.NewDecoder(req.Body).Decode(&student)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		studentsGroup[student.Id] = student
		fmt.Fprintf(res, "{\n    \"message\": \"student registered\"\n}")
	}
}

func studentsById(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method, req.URL.Path)
	res.Header().Set("Content-Type", "application/json")

	id, err := strconv.ParseUint(strings.TrimPrefix(req.URL.Path, "/students/"), 10, 64)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	student, isValid := studentsGroup[id]
	if !isValid {
		http.Error(res, "{\n    \"error\": \"student not found\"\n}", http.StatusNotFound)
		return
	}

	switch req.Method {
	case "GET":
		studentJson, err := json.MarshalIndent(student, "", "    ")
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		if err = req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
			return
		}

		fmt.Fprintf(res, string(studentJson))
	case "DELETE":
		delete(studentsGroup, id)
		fmt.Fprintf(res, "{\n    \"message\": \"student deleted\"\n}")
	case "PUT":
		subject := req.FormValue("subject")
		grade, err := strconv.ParseFloat(req.FormValue("grade"), 64)
		if err != nil {
			http.Error(res, "{\n    \"error\": \"invalid grade\"\n}", http.StatusBadRequest)
			return
		}

		student.Subjects[subject] = grade
		fmt.Fprintf(res, "{\n    \"message\": \"grade updated\"\n}")
	}
}

func main() {
	studentsGroup = make(map[uint64]Student)
	nextId = 1

	http.HandleFunc("/students", students)
	http.HandleFunc("/students/", studentsById)

	fmt.Println("Server running...")
	fmt.Println("http://localhost:9000/")
	http.ListenAndServe(":9000", nil)
}
