package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}

	sum := 0

	for _, val := range arr {
		sum += val
	}

	for index, val := range arr {
		fmt.Printf("Index %d value %d \n", index, val)
	}

	// fmt.Println(sum)

}
