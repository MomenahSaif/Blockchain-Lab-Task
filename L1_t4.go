package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

type Student struct {
	rollnumber int
	name       string
	address    string
	subjects   []string // Array of subjects a student is studying
}

func NewStudent(rollno int, name, address string, subjects []string) *Student {
	s := new(Student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	s.subjects = subjects
	return s
}

type StudentList struct {
	list []Student
}

func (ls *StudentList) CreateStudent(rollno int, name, address string, subjects []string) *Student {
	st := *NewStudent(rollno, name, address, subjects)
	ls.list = append(ls.list, st)
	return &st
}

func calculateHash(student Student) string {
	data := fmt.Sprintf("%d%s%s%v", student.rollnumber, student.name, student.address, student.subjects)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func (ls *StudentList) Print() {
	for i, student := range ls.list {
		fmt.Printf("========================= List %d %s\n", i, strings.Repeat("=", 25))
		// Print student details
		fmt.Printf("student rollno   %d\n", student.rollnumber)
		fmt.Printf("student name     %s\n", student.name)
		fmt.Printf("student address  %s\n", student.address)
		fmt.Printf("student subjects %v\n", student.subjects)

		// Calculate and print hash for the block
		hash := calculateHash(student)
		fmt.Printf("Block Hash       %s\n", hash)

	}
}

func main() {
	studentList := new(StudentList)
	studentList.CreateStudent(24, "Asim", "AAAAAA", []string{"Math", "Physics", "Chemistry"})
	studentList.CreateStudent(25, "Naveed", "BBBBBB", []string{"Biology", "English", "Computer Science"})

	// Print the list of students
	studentList.Print()
}
