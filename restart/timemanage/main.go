package main

import (
	"fmt"
	"time"
)

func main() {
	currTime := time.Now()
	fmt.Println(currTime)
	fmt.Print(currTime.Format("01-02-2006 Monday"))
}
