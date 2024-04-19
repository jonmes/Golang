package main

import "fmt"

// Declare a channel to store data fetched from the database
// var results = make(chan string)

func main() {
	// Declare a channel to store data fetched from the database
	var c = make(chan int)

	go process(c)

	// var i = <- c
	for j := range c {
		fmt.Println(j)
	}

}

func process(c chan int) {
	for i := 0; i < 5; i++ {
		// fmt.Println(i)
		c <- i
	}
	defer close(c)
}
