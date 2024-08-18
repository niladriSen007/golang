package main

import "fmt"

func main() {
	/* name := "World"
	fmt.Println("Hello,", name) */

	//classic for loop
	/* for i := 0; i < 10; i++ {
		fmt.Println(i)
	} */

	//while loop
	/* 	i := 0
	   	for i < 10 {
	   		fmt.Println(i)
	   		i++
	   	} */

	//infinite loop
	/* for {
		fmt.Println("Hello")
	} */

	/* for i := range 3 {
		fmt.Println(i)
	} */

	/*
		switch case */
	/* whoAmI := func(val interface{}) {
		switch t := val.(type) {
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		default:
			fmt.Println("other", t)
		}
	}

	whoAmI(18.8) */

	/* var nums [4]string */
	/* nums := [4]string{"zero", "one", "two", "three"}
	nums[0] = "one"
	nums[1] = "two"

	fmt.Println(len(nums))
	fmt.Println(nums) */

	//2d array
	nums := [2][3]int{{2, 3, 4}, {1, 2, 3}}
	fmt.Println(nums)

}
