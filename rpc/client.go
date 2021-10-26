package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", ":9999")
	if err != nil {
		log.Fatal("Error:", err)
	}

	exit := false
	for !exit {
		fmt.Print("\033[H\033[2J1. Register Student Grade\n2. Get Student GPA\n3. Get Students Average GPA\n4. Get Subject Grades Average\n5. Exit\n\nOption: ")
		var option string
		fmt.Scanln(&option)

		switch option {
		case "1":
			student := make(map[string]string)
			var studentName, subject, grade, status string

			fmt.Print("Name: ")
			fmt.Scanln(&studentName)
			fmt.Print("Subject: ")
			fmt.Scanln(&subject)
			fmt.Print("Grade: ")
			fmt.Scanln(&grade)

			student["name"] = studentName
			student["subject"] = subject
			student["grade"] = grade

			err = client.Call("Teacher.RegisterStudentGrade", &student, &status)
			if err != nil {
				fmt.Println("Error:", err)
				break
			}
			fmt.Println("Status:", status)
		case "2":
			var student string
			var studentGPA float64

			fmt.Print("Name: ")
			fmt.Scanln(&student)

			err = client.Call("Teacher.GetStudentGPA", &student, &studentGPA)
			if err != nil {
				fmt.Println("Error:", err)
				break
			}
			fmt.Println(student, "GPA:", studentGPA)
		case "3":
			var studentsAverageGPA float64

			err = client.Call("Teacher.GetStudentsAverageGPA", 0, &studentsAverageGPA)
			if err != nil {
				fmt.Println("Error:", err)
				break
			}
			fmt.Println("Students average GPA:", studentsAverageGPA)
		case "4":
			var subject string
			var subjectGradesAverage float64

			fmt.Print("Subject: ")
			fmt.Scanln(&subject)

			err = client.Call("Teacher.GetSubjectGradesAverage", &subject, &subjectGradesAverage)
			if err != nil {
				fmt.Println("Error:", err)
				break
			}
			fmt.Println(subject, "average grade:", subjectGradesAverage)
		case "5":
			fmt.Println("Closing client connection gracefully...")
			exit = true
			err = client.Close()
			if err != nil {
				log.Fatal("Error: ", err)
			}
		default:
			fmt.Println("Invalid option")
		}

		fmt.Print("Press enter to continue...")
		fmt.Scanln(&option)
	}
}
