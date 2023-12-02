package main

import "fmt"

func retString() string {
	return "Niladri"
}

func retInt() int {
	return 3
}

func retTwo(name string, age int) (int, string) {
	return age, name
}
func retTwoNew(age int) int {
	return age
}
func retTwoNew1(age int) (age1 int) {
	age1 = age + 1
	return
}
func retTwoShort(name string, age int) (age1 int, name1 string) {
	age1 = age + 1
	name1 = name
	return
}

type Person struct {
	name string
	age  int
	role string
}

func main() {
	// fmt.Print("Nil")
	cards := []string{"Nil", "Niladri"}
	// fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, "=>", card)
	}

	var newPerson Person
	newPerson.name = "Niladri"
	newPerson.age = 25
	newPerson.role = "Developer"
}
