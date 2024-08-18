package fileops

import (
	"errors"
	"fmt"
	"os"
)

func WriteIntoFile(age float64) {
	myAgeText := fmt.Sprint(age)
	os.WriteFile("myAge.txt", []byte(myAgeText), 0644)
}

func CreateFile() {
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

func ReadFile() (float64, error) {
	fmt.Println("Reading a file...")
	filename := "myAge.txt"

	data, err := os.ReadFile(filename)

	if err != nil {
		err = errors.New("error reading the file")
		return 1000, err
	}

	/* fmt.Println("Data read from the file:", string(data)) */
	fmt.Printf("Data read from the file: %s \n", data)

	return 0, nil
}
