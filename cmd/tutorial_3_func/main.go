package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe("print function!")
	var result, remainder, err = intDivision(10, 0)

	// if else statements

	// if err != nil {
	// 	fmt.Printf(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("The result of the integer division is %d\n", result)
	// } else {
	// 	fmt.Printf("result: %v, remainder: %d\n", result, remainder)
	// }

	// switch statement

	switch {
	case err != nil:
		fmt.Printf("Error: %v\n", err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %d\n", result)
	default:
		fmt.Printf("result: %v, remainder: %d\n", result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Printf("result: %v, remainder: %d\n", result, remainder)
	case 1, 2, 3, 4, 5, 6:
		fmt.Printf("The result of the integration division is %d", result)
	default:
		fmt.Printf("show error, %v\n", err.Error())
	}
}

// This function takes a string as an argument and prints it; this function doesn't return any value
func printMe(printValue string) {
	fmt.Println(printValue)
}

// This function takes two integers as arguments and returns their sum and remainder
func intDivision(a int, b int) (int, int, error) {

	var err error
	if b == 0 {
		err = errors.New("division by zero is not allowed ")
		return 0, 0, err
	}
	var result int = a / b
	var remainder int = a % b
	return result, remainder, err
}
