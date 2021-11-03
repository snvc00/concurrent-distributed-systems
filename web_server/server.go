package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/rpc"
)

func loadHTML(filename string) string {
	html, _ := ioutil.ReadFile(filename)

	return string(html)
}

func getServiceClient() *rpc.Client {
	client, err := rpc.DialHTTP("tcp", ":9999")
	if err != nil {
		log.Fatal("Error:", err)
	}

	return client
}

func registerGrade(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method, req.URL.Path)
	res.Header().Set("Content-Type", "text/html")

	switch req.Method {
	case "GET":
		fmt.Fprintf(res, loadHTML("register-grade.html"), "")
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
			return
		}

		var status string
		client := getServiceClient()
		student := make(map[string]string)
		student["name"] = req.FormValue("student")
		student["subject"] = req.FormValue("subject")
		student["grade"] = req.FormValue("grade")

		err := client.Call("Teacher.RegisterStudentGrade", &student, &status)
		if err != nil {
			fmt.Fprintf(res, "RPC call error %v", err)
			return
		}

		fmt.Fprintf(res, loadHTML("register-grade.html"), status)
	}
}

func studentGPA(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method, req.URL.Path)
	res.Header().Set("Content-Type", "text/html")

	switch req.Method {
	case "GET":
		fmt.Fprintf(res, loadHTML("student-gpa.html"), "", 0.0)
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
			return
		}

		var studentGPA float64
		client := getServiceClient()
		student := req.FormValue("student")

		err := client.Call("Teacher.GetStudentGPA", &student, &studentGPA)
		if err != nil {
			fmt.Fprintf(res, "RPC call error %v", err)
			return
		}

		fmt.Fprintf(res, loadHTML("student-gpa.html"), student, studentGPA)
	}
}

func subjectGPA(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method, req.URL.Path)
	res.Header().Set("Content-Type", "text/html")

	switch req.Method {
	case "GET":
		fmt.Fprintf(res, loadHTML("subject-gpa.html"), "", 0.0)
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
			return
		}

		var subjectGPA float64
		client := getServiceClient()
		subject := req.FormValue("subject")

		err := client.Call("Teacher.GetSubjectGradesAverage", &subject, &subjectGPA)
		if err != nil {
			fmt.Fprintf(res, "RPC call error %v", err)
			return
		}

		fmt.Fprintf(res, loadHTML("subject-gpa.html"), subject, subjectGPA)
	}
}

func studentsAverageGPA(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method, req.URL.Path)
	switch req.Method {
	case "GET":
		var studentsAverageGPA float64
		client := getServiceClient()

		err := client.Call("Teacher.GetStudentsAverageGPA", 0, &studentsAverageGPA)
		if err != nil {
			fmt.Fprintf(res, "RPC call error %v", err)
			return
		}

		res.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(res, loadHTML("students-average-gpa.html"), studentsAverageGPA)
	}
}

func main() {
	fmt.Println("Server starting...")

	http.HandleFunc("/register-grade", registerGrade)
	http.HandleFunc("/student-gpa", studentGPA)
	http.HandleFunc("/subject-gpa", subjectGPA)
	http.HandleFunc("/students-average-gpa", studentsAverageGPA)

	fmt.Println("Server running...")
	fmt.Println("http://localhost:9000/register-grade")
	http.ListenAndServe(":9000", nil)
}
