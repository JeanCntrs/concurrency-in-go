package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Processes start")
	start := time.Now()
	doSomething()
	doSomethingElse()

	fmt.Println("I guess I'm done")
	elapsed := time.Since(start)
	fmt.Printf("Processes end, took %s", elapsed)
}

func doSomething() {
	time.Sleep(time.Second * 2)
	fmt.Println("I've done something")
}

func doSomethingElse() {
	time.Sleep(time.Second * 2)
	fmt.Println("I've done something else")
}
