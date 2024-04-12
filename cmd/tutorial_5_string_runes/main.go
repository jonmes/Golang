package main

import "fmt"

func main() {
	var myString = []rune("Résumé")
	var indexedString = myString[0]
	fmt.Printf("Hello World! %v", indexedString)
	fmt.Printf("\nHello World! %v\n", len(myString))
	for i, v := range myString {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}
	var strSlice = []string{"s", "u", "m", "m", "e", "r"}
	var catStr = ""
	for i := range strSlice {
		catStr += strSlice[i]
	}
	fmt.Printf("Concatenated String: %v\n", catStr)
}
