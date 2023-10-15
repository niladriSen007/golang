// package main

// import "fmt"

// // func getCalc(x, y int) (int, int, int, int) {
// // 	return (x + y), (x - y), (x * y), (x / y)
// // }

// var addition int = 0
// var sub int = 0
// var mul int = 2
// var div int = 1

// func getCalcNew(args ...int) (addition, mul int) {
// 	for _, val := range args {
// 		addition += val
// 		// sub -= val
// 		mul *= val
// 		// div /= val
// 		fmt.Print(mul)

// 	}
// 	return addition, mul
// }

// func main() {
// 	// sum, sub, mul, div := getCalc(5, 6)

// 	addition, mul := getCalcNew(5, 6)

// 	// fmt.Printf("Sum is %d\n Subtraction is %d\n Multiplication is %d\n Division is %d\n ", sum, sub, mul, div)
// 	fmt.Printf("Sum is %d\n  Multiplication is %d\n  ", addition, mul)
// }
