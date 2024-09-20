package main

import (
	"fmt"
	"strings"
)

type Student struct {
	rollnumber int
	name       string
	address    string
}

func NewStudent(rollno int, name, address string) *Student {
	s := new(Student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	return s
}

type StudentList struct {
	list []Student
}

func (ls *StudentList) CreateStudent(rollno int, name, address string) *Student {
	st := *NewStudent(rollno, name, address)
	ls.list = append(ls.list, st)
	return &st
}

func (ls *StudentList) Print() {
	for i, student := range ls.list {
		fmt.Printf("========================= List %d %s\n", i, strings.Repeat("=", 25))
		fmt.Printf("student rollno   %d\n", student.rollnumber)
		fmt.Printf("student name     %s\n", student.name)
		fmt.Printf("student address  %s\n", student.address)
	}
}

func main() {
	studentList := new(StudentList)
	studentList.CreateStudent(24, "Asim", "AAAAAA")
	studentList.CreateStudent(25, "Naveed", "BBBBBB")

	// Print the list of students
	studentList.Print()
}
