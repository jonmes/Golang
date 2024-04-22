package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Declare a channel to store data fetched from the database
var MAX_CHICKEN_PRICE float32 = 10

func main() {
	// Declare a channel to store data fetched from the database
	// var c = make(chan int)

	// go process(c)

	// This is a synchronous operation that will block the main goroutine until the channel is closed.
	// This loop will run until the channel is closed
	// for j := range c {
	// 	fmt.Println(j)
	// }

	var chickenChannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
	}
	sendMessage(chickenChannel)

}

// func process(c chan int) {
// 	for i := 0; i < 50000; i++ {
// 		c <- i
// 	}
// 	defer close(c)
// }

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 3)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string) {
	fmt.Printf("The chicken is available at %s\n", <-chickenChannel)
}
