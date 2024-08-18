package main

import "example.com/structs/personpointer"

func main() {

	//Creating a Person object by passing values in the order of the struct fields
	/* var p1 Person */
	/* p1 := Person{"John", 25, "john@emp.com", time.Now()} */

	/* p1.printPersonActual()
	p1.clearUsername()
	p1.printPersonActual() */

	/* var personPointer *Person */
	personPointer, err := personpointer.NewPerson("John", 25, "john@emp.com")

	if err != nil {
		panic(err)
	}

	personPointer.PrintPersonActual()
	personPointer.ClearUsername()
	personPointer.PrintPersonActual()

	/* printPerson(p1) */

	/* printPersonByPointer(&p1) */

	/* p1.printPerson() */

}
