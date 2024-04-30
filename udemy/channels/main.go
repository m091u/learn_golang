package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
    ch := make(chan int, 5)
	wg.Add(2)

	//receiving goroutine from the channel
	go func(ch <-chan int){  //channel is send only by adding the paremeter
		// i:= <- ch
		for i := range ch {   // range is used to process more incoming data from the sending routines
			fmt.Println("Received", i)
		}
		wg.Done()
	}(ch)

	//sending goroutine to the channel
	go func(ch chan<- int){  //channel is receive only as we pass the specific parameter
		ch <- 42
		ch <- 34    //buffer scenario where more data is being sent
		close(ch)   // needs to be added to let the receiver that the channel is no longer sending data
		wg.Done()
	}(ch)

	wg.Wait()
}