package main

import (
	"fmt"
	"sync"
)

type Recipient struct {
	Name  string
	Email string
}

func main() {
	var wg sync.WaitGroup
	recipientChanel := make(chan Recipient)

	fmt.Println("Welcome to email dispatcher! ")
	go func() {
		err := loadRecipient("./emails.csv", recipientChanel)
		if err != nil {
			fmt.Println(err)
		}
	}()
	workerCount := 5
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go emailWorker(i, recipientChanel, &wg)
	}
	wg.Wait()
}
