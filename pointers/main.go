package main

import "fmt"

func getOld(age int) int {
	return age - 18
}

/* func getOldPointer(agePointer *int) int {
	return *agePointer - 18
} */
func getOldPointer(agePointer *int) {
	*agePointer = *agePointer - 18
}

func main() {
	age := 25
	/* oldAge := getOld(age)
	fmt.Println("My old age is", oldAge) */

	var agePointer *int
	agePointer = &age

	/* 	fmt.Println("Address of age:", agePointer)
	   	fmt.Println("Value of age:", *agePointer) */

	/* oldAgePointer := getOldPointer(agePointer) */
	getOldPointer(agePointer)
	fmt.Println("My current age is", age)
	/* fmt.Println("I am", oldAgePointer, "years adult") */

}
