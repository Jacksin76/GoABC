package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	// Create channels
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)

	go func() {
		defer wg.Done()
		for i := 0; i <= 10; i++ {
			fmt.Println("A", <-chA) // Receive data from channel chA
			time.Sleep(time.Millisecond * 100)
			chB <- i // Send data to channel chB
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i <= 10; i++ {
			fmt.Println("B", <-chB) // Receive data from channel chB
			time.Sleep(time.Millisecond * 100)
			chC <- i // Send data to channel chC
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i <= 10; i++ {
			chA <- i                // Send data to channel chA
			fmt.Println("C", <-chC) // Receive data from channel chC
			time.Sleep(time.Millisecond * 100)
		}
	}()

	wg.Wait()
	// Close the channels
	close(chA)
	close(chB)
	close(chC)
}
