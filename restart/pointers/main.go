package main

import "fmt"

func main() {
	fmt.Println(`Welcome to Go Pointers`)

	num := 12
	var ptr = &num
	fmt.Println(ptr)
	fmt.Println(*ptr)
}
