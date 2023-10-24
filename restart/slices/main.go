package main

import "fmt"

func main() {
	var fruits = []string{"apple", "banana"}
	fruits = append(fruits, "mango")
	fruits = append(fruits[1:])
	fmt.Println(fruits)

	// fruitsNew := make([]string, 4)

	if num := 14; num < 10 {
		fmt.Println("Num less than 10")
	} else {
		fmt.Println("Num greater than 10")
	}

}
