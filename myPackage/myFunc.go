package myFunc

import "fmt"

func PrintMessage(message string) {
	fmt.Println(message)
}

func TcsHike(Band string) {
	fmt.Printf("Your Band is: %T\n", Band)
	switch Band {
	case "A":
		fmt.Println("7%")
	case "B":
		fmt.Println("4%")
	case "C":
		fmt.Println("2%")
	default:
		fmt.Println("Employees have to pay TCS")
	}
}
