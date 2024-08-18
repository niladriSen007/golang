package main

import (
	"fmt"
	"os"
)

func writeIntoFile(age float64) {
	myAgeText := fmt.Sprint(age)
	os.WriteFile("myAge.txt", []byte(myAgeText), 0644)
}

func createFile() {
	fmt.Println("Creating a file...")
	file, err := os.Create("myAge.txt")

	if err != nil {
		panic(err)
	}
	length, err := file.WriteString("Hello world")

	if err != nil {
		panic(err)
	}

	fmt.Printf("File name %s", file.Name())
	fmt.Println("Length of the file", length)
}

func readFile() {
	fmt.Println("Reading a file...")
	filename := "myAge.txt"

	data, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	/* fmt.Println("Data read from the file:", string(data)) */
	fmt.Printf("Data read from the file: %s", data)
}

func main() {
	const myAge = 24
	/* writeIntoFile(myAge) */
	createFile()
	readFile()
}
