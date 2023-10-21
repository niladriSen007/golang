package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Hello welcome"
	fmt.Println(welcome)
	input := bufio.NewReader(os.Stdin)
	fmt.Println(("Enter a number"))

	io, _ := input.ReadString('\n')
	fmt.Println(io)

}
