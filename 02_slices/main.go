package main

import "fmt"

func main() {
	/* var goSlices []int */
	goSlices := make([]bool, 5, 10)

	goSlices = append(goSlices, true)
	 
	fmt.Println(goSlices)
	fmt.Println(len(goSlices))
	fmt.Println(cap(goSlices))

}
