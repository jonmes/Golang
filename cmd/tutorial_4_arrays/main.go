package main

import (
	"fmt"
	"time"
)

func main() {
	// intArr := [4]int32{1, 2, 3, 4} for an array of size 4
	intArr := [...]int32{1, 2, 3, 4, 4, 5, 6, 7, 8, 9, 9, 99} // for an unknown size array
	intArr[10] = 123

	fmt.Println(intArr[0])
	fmt.Println(intArr[1])
	fmt.Println(intArr[2])
	fmt.Println(intArr[10])

	// =============== Slice ===============
	var intSlice []int32 = []int32{3, 4, 5}

	fmt.Println(intSlice[0])

	intSlice = append(intSlice, 23444, 33, 444, 888, 88, 33)

	fmt.Println(cap(intSlice))
	fmt.Println(intSlice[cap(intSlice)-5]) // Cant access the cap length element of the slice ** By default the cap is 2 times the length of the slice **

	// var intSlice2 []int32 = make([]int32, 3, 8)
	var intSlice2 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice2[2])
	// make([]int32, 3, 10) // 3 is the length of the slice and 10 is the capacity of the slice

	// =============== Map ===============

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Eve": 22, "Abel": 33}
	fmt.Println(myMap2["Adam"])
	var age, ok = myMap2["Eve"]

	fmt.Println(age, ok)

	// delete Map value

	delete(myMap2, "Adam")

	// =============== Loops ===============

	for name, age := range myMap2 {
		fmt.Printf("here me out name %v: and age : %v", name, age)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var n int = 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v\n", timeLoop(testSlice2, n))

}

func timeLoop(testSlice []int, n int) time.Duration {
	start := time.Now()
	for i := 0; i < n; i++ {
		testSlice = append(testSlice, i)
	}
	return time.Since(start)
}
