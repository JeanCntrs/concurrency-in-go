package main

import (
	"fmt"
	"time"
)

var ch = make(chan string)

func main() {
	fmt.Println("Processes start")

	start := time.Now()
	go doSomething()
	go doSomethingElse()

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("I guess I'm done")
	elapsed := time.Since(start)
	fmt.Printf("Processes end, took %s", elapsed)
}

func doSomething() {
	time.Sleep(time.Second * 2)
	fmt.Println("I've done something")
	ch <- "doSomething finished"
}

func doSomethingElse() {
	time.Sleep(time.Second * 2)
	fmt.Println("I've done something else")
	ch <- "doSomethingElse finished"
}
