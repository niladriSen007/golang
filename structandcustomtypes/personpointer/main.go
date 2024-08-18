package personpointer

import (
	"errors"
	"fmt"
	"time"
)

type Person struct {
	username  string
	age       int
	email     string
	createdAt time.Time
}

func NewPerson(username string, age int, email string) (*Person, error) {
	if username == "" || age == 0 || email == "" {
		return nil, errors.New("username cannot be empty")
	}
	return &Person{username,
		age,
		email,
		time.Now()}, nil
}

func PrintPerson(p Person) {
	fmt.Println("Name: ", p.username)
	fmt.Println("Age: ", p.age)
	fmt.Println("Email: ", p.email)
	fmt.Println("Created At: ", p.createdAt)
}

func PrintPersonByPointer(p *Person) {
	fmt.Println("Printing Person by Pointer")
	fmt.Println("Name: ", p.username)
	fmt.Println("Age: ", p.age)
	fmt.Println("Email: ", p.email)
	fmt.Println("Created At: ", p.createdAt)
}

// This is a method of the Person struct
func (p Person) PrintPerson() {
	fmt.Println("Name: ", p.username)
	fmt.Println("Age: ", p.age)
	fmt.Println("Email: ", p.email)
	fmt.Println("Created At: ", p.createdAt)
}

// This is a method of the Person struct
func (p *Person) PrintPersonActual() {
	fmt.Println("Name: ", p.username)
	fmt.Println("Age: ", p.age)
	fmt.Println("Email: ", p.email)
	fmt.Println("Created At: ", p.createdAt)
}

func (p *Person) ClearUsername() {
	p.username = ""
}
