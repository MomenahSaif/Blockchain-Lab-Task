package main

import "fmt"

type employee struct {
	name     string
	salary   int
	position string
}

type company struct {
	companyName string
	employees   []employee
}

func printCompanyDetails(comp company) {
	fmt.Println("Company Name:", comp.companyName)
	fmt.Println("Employees:")
	for _, emp := range comp.employees {
		fmt.Println(" - Name:", emp.name, "| Salary:", emp.salary, "| Position:", emp.position)
	}
}

func main() {
	// Create three employee records
	emp1 := employee{"Amir", 80000, "Full-Stack Developer"}
	emp2 := employee{"Sara", 75000, "Backend Developer"}
	emp3 := employee{"John", 90000, "DevOps Engineer"}

	// Create an array of employees
	emplys := []employee{emp1, emp2, emp3}

	// Create a company struct and add values to it
	myCompany := company{"Tetra", emplys}

	// Print the company details
	printCompanyDetails(myCompany)
}
