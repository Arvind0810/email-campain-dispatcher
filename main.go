package main

import (
	"fmt"
	"sync"
)

/* func EvenOdd(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	val := <-ch
	if val%2 == 0 {
		fmt.Printf("I am go routine %d, and I am Even\n", val)
	} else {
		fmt.Printf("I am go routine %d, and I am Odd\n", val)
	}
}
func main() {
	var wg sync.WaitGroup
	chnl := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		chnl <- i
		go EvenOdd(chnl, &wg)
	}

	wg.Wait()
} */

// Two goroutines print 1..n in strict order by passing a "turn" token back
// and forth over unbuffered channels (ping-pong).
func main() {
	var wg sync.WaitGroup
	// Unbuffered and carrying struct{} (zero bytes): the channels are pure
	// signals, and each send blocks until the other goroutine is ready to
	// receive, which is what enforces the alternation.
	oddChan := make(chan struct{})
	evenChan := make(chan struct{})

	n := 10

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i < n; i += 2 {
			<-oddChan
			fmt.Printf("I am go routine %d, and I am Odd\n", i)
			evenChan <- struct{}{}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 2; i <= n; i += 2 {
			<-evenChan
			fmt.Printf("I am go routine %d, and I am even\n", i)
			// On the last iteration (i == n) the odd goroutine has already
			// exited, so this send has no receiver and blocks forever; with
			// main stuck in wg.Wait(), the runtime reports a deadlock.
			if i < n {
				oddChan <- struct{}{}
			}
		}
	}()

	// Neither goroutine can start until someone hands out the first turn.
	oddChan <- struct{}{}
	wg.Wait()
	// close(oddChan)
	// close(evenChan)
}
