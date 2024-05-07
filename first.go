package main

import (
	"bufio"
	"fmt"
	myFunc "golangpractice/myPackage"
	"os"
)

func main() {
	myFunc.PrintMessage("***** GoLannng is a hell.😥 I'm Dead Bro *****")
	myFunc.TcsHike("A")
	reader := bufio.NewReader(os.Stdin)
	yourBand, _ := reader.ReadString('\n')
	fmt.Printf("Type of yourBand: %T\n", yourBand)
	fmt.Println("Your Band is: ", yourBand)
	myFunc.TcsHike(yourBand)
}
