package main

import "fmt"

func main() {
	const rate = 9

	var investment float64

	fmt.Print("Investment: ")
	fmt.Scan(&investment)

	investment = investment + (investment * rate / 100)
	fmt.Println("Investment after 1 year:", investment)
}
