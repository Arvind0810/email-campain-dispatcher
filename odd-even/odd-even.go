package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	oddChan := make(chan int)
	evenChan := make(chan int)

	wg.Add(2)

	go oddFunc(oddChan, evenChan, &wg)
	go evenFunc(evenChan, oddChan, &wg)

	oddChan <- 0

	wg.Wait()
}

func oddFunc(odd chan int, even chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		<-odd
		fmt.Printf("I am %d go routine, and I am odd.", i)
		even <- i
	}
}
func evenFunc(even chan int, odd chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		<-even
		fmt.Printf("I am %d go routine, and I am odd.", i)
		odd <- i
	}
}
