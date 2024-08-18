package main

import "fmt"

func main() {
	slices := make([]int, 5, 10)

	for _, element := range slices {
		fmt.Println(element)
	}
}
