package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50) // With 50 buffer
	wg.Add(2)
	go func(ch <-chan int) { // Receive only channel
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) { // Send only channel
		ch <- 42
		ch <- 47
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()

}
