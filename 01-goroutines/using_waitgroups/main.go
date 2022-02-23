package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{} // This creates our waitgroup called, "wg".

func main() {
	fmt.Println("Processes start")

	wg.Add(2) // This sets our wait group to 2, meaning we will wait for 2 goroutines to complete.
	start := time.Now()
	go doSomething()     // Adding the keyword "go" creates a goroutine.
	go doSomethingElse() // Adding the keyword "go" creates a goroutine.
	wg.Wait()            // This blocks (waits) until wg=0 meaning all goroutines are done.

	fmt.Println("I guess I'm done")
	elapsed := time.Since(start)
	fmt.Printf("Processes end, took %s", elapsed)
}

func doSomething() {
	time.Sleep(time.Second * 2)
	fmt.Println("I've done something")
	wg.Done() // This decrements wg by one, indicating that doSomething is done.
}

func doSomethingElse() {
	time.Sleep(time.Second * 2)
	fmt.Println("I've done something else")
	wg.Done() // This decrements wg by one, indicating that doSomethingElse is done.
}
