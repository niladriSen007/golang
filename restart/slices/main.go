package main

import "fmt"

func main() {
	var fruits = []string{"apple", "banana"}
	fruits = append(fruits, "mango")
	fruits = append(fruits[1:])
	fmt.Println(fruits)

	// fruitsNew := make([]string, 4)

}
