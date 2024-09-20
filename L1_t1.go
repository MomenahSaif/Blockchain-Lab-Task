// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func printstruct(pers Person) {
	fmt.Println(pers.name)
	fmt.Println(pers.age)
	fmt.Println(pers.job)
	fmt.Println(pers.salary)

}

func main() {
	var pers1 Person
	var pers2 Person

	// Pers1 specification
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	// Pers2 specification
	pers2.name = "Cecilie"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary = 4500

	fmt.Println("Person 1")
	printstruct(pers1)
	fmt.Println("Person 2")
	printstruct(pers2)
}
