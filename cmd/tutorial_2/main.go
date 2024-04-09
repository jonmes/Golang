package main

import "fmt"
import "unicode/utf8"

func main() {
	var intNum int = 10
	intNum = intNum / 3
	fmt.Println("intNum:", intNum)

	var myString string = "Hello, world!"

	fmt.Println(len(myString))                    // Characters outside the ASCII range are counted as multiple characters or take multiple bytes
	fmt.Println(utf8.RuneCountInString(myString)) // This is the correct way to count the number of characters in a string

	var myRune rune = 'a' // A rune is an alias for int32 and represents a Unicode code point in Go (UTF-8 encoded), the are like characters in other languages
	fmt.Println(myRune)

	var myBool bool = true
	fmt.Println(myBool)

	var1, var2 := 1, 2 // Initialize multiple variables with walrus operator
	fmt.Println(var1, var2)

	const pi float32 = 3.1415 // Constants are immutable

	var1 = 44
	// pi = "new value" // This will cause an error because constants are immutable
	fmt.Println(var1, pi)

	// Functions

	printMe()
}

func printMe() {
	fmt.Println("Hello from printMe()!")
}
