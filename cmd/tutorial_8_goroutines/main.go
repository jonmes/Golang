package main

import (
	"fmt"
	"sync"
	"time"
)

// Declare a mutex variable to protect shared data access between goroutines
var m = sync.Mutex{}

// Declare a WaitGroup variable to wait for goroutines to finish
var wg = sync.WaitGroup{}

// Data to be fetched from the database
var dbData = []string{"data1", "data2", "data3", "data4", "data5"}

// Declare a slice to store data fetched from the database
var results = []string{}

func main() {
	// Start a goroutine to fetch data from the database
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}

	// Wait for the goroutines to finish
	wg.Wait()
	fmt.Println("All database calls completed", time.Since(t0))
	fmt.Println("Results are", results)
}

// This function emulates database query

func dbCall(i int) {
	// Simulate a database call
	fmt.Println("Database call for", i)
	var delay float32 = 5000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	m.Lock()
	results = append(results, dbData[i])
	m.Unlock()
	wg.Done()
	fmt.Println("Database call for", dbData[i], "completed")
}
