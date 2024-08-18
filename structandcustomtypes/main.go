package main

import (
	"fmt"
	"time"
)

type Person struct {
	username  string
	age       int
	email     string
	createdAt time.Time
}

func printPerson(p Person) {
	fmt.Println("Name: ", p.username)
	fmt.Println("Age: ", p.age)
	fmt.Println("Email: ", p.email)
	fmt.Println("Created At: ", p.createdAt)
}

func printPersonByPointer(p *Person) {
	fmt.Println("Printing Person by Pointer")
	fmt.Println("Name: ", p.username)
	fmt.Println("Age: ", p.age)
	fmt.Println("Email: ", p.email)
	fmt.Println("Created At: ", p.createdAt)
}

// This is a method of the Person struct
func (p Person) printPerson() {
	fmt.Println("Name: ", p.username)
	fmt.Println("Age: ", p.age)
	fmt.Println("Email: ", p.email)
	fmt.Println("Created At: ", p.createdAt)
}

// This is a method of the Person struct
func (p *Person) printPersonActual() {
	fmt.Println("Name: ", p.username)
	fmt.Println("Age: ", p.age)
	fmt.Println("Email: ", p.email)
	fmt.Println("Created At: ", p.createdAt)
}

func (p *Person) clearUsername() {
	p.username = ""
}

func main() {

	p1 := Person{"John", 25, "john@emp.com", time.Now()}

	/* printPerson(p1) */

	/* printPersonByPointer(&p1) */

	/* p1.printPerson() */

	p1.printPersonActual()
	p1.clearUsername()
	p1.printPersonActual()
}
