package main

import "fmt"

func main() {
	// Declare a pointer variable
	var ptr *int

	// Declare a variable
	var num int = 10

	// Assign the address of num to ptr
	ptr = &num

	// Display the value of num
	fmt.Println("Value of num =", num)

	// Display the value of num using ptr
	fmt.Println("Value of num using ptr =", *ptr)

	fmt.Println("Address of num =", &num)

	// Display the address of num using ptr
	fmt.Println("Address of num using ptr =", ptr)

	// Change the value of num using ptr
	*ptr = 20

	// Display the value of num
	fmt.Println("Value of num =", num)

	// Display nill pointer
	var ptr2 *int
	// fmt.Println("Value of ptr2 =", *ptr2) 	// error: invalid memory address or nil pointer dereference
	fmt.Println("Value of ptr2 =", ptr2)
}
