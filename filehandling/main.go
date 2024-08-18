package main

import (
	"example.com/fileops"
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

func main() {
	/* const myAge = 24 */
	/* writeIntoFile(myAge) */
	fileops.CreateFile()
	fileops.ReadFile()
	fmt.Println(randomdata.SillyName())
}
