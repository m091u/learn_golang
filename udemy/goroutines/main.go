package main

import (
	"fmt"
	// "time"
	"sync"
)

var wg = sync.WaitGroup{}
func main() {
	go sayHello() //go creates a green thread, creates an abstraction of a thread
	// time.Sleep(10 * time.Millisecond)
	// select{}

	wg.Add(1)
	msg := "help"
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	wg.Wait()
	// time.Sleep(10 * time.Millisecond)
}

func sayHello() {
	fmt.Println("Hello, World!")
}
