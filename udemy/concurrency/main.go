package main

import (
	"fmt"
	"sync"
	"time"
)

func pinger(pinger <-chan int, ponger chan<- int) {
	for {
		<-pinger
		fmt.Println("ping")
		time.Sleep(time.Second)
		ponger <- 1
	}
}

func ponger(pinger chan <- int, ponger <-chan int) {
	for {
		<-ponger
		fmt.Println("pong")
		time.Sleep(time.Second)
		pinger <- 1
	}
}
// func heavy() {
// 	time.Sleep(time.Second * 5)
// 	fmt.Println("heavy task completed")
// }

func main() {
	// using select to wait indefinitely  and not close the main thread
	// go heavy()
	// fmt.Println("main thread completed")
	// select {}

	//wait groups that offers ADD, DONE & WAIT functionalities
	wg := &sync.WaitGroup{}

	wg.Add(2) //add functions to be completed before clasing main thread
	//anonimous function with no name. it is called () after the {}
	go func() {
		fmt.Println("hello")
		wg.Done()
	}()
	go func() {
		fmt.Println("world")
		wg.Done() //task to be completed
	}()

	wg.Wait() //wait for functions to be executed
	fmt.Println("main thread completed")

	// channels
	c := make(chan int)

	//send in a go routine
	go func() {
		c <- 1
	}()

	//sniff
	val := <-c
	fmt.Println(val)

	//send in a go routine
	go func() {
		c <- 2
	}()

	time.Sleep(time.Second * 2)
	val = <-c
	fmt.Println(val)

	fmt.Println(c)

	//making buffer channels that allow passing more data
	bc := make(chan int, 3)

	go func() {
		bc <- 1
		bc <- 2
		bc <- 3
		bc <- 4
		close(bc)
	}()

	for i := range bc {
		fmt.Println(i)
	}



	//ping pong game
	ping := make(chan int)
	pong := make(chan int)

	go pinger(ping, pong)
	go ponger(ping, pong)

	//routine strats by sending into the ping channel
	ping <- 1
	select{}
}
