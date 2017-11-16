package main

import (
	"fmt"
	"sync"
	"time"
)

var group sync.WaitGroup

func main() {
	done := make(chan bool)
	go hardTask(done)

	wg.Add(1)
	go loader(done)

	wg.Wait()
}

func hardTask(done chan<- bool) {
	fmt.Println("Start task")
	time.Sleep(time.Second * 1)
	done <- true
}

func loader(done <-chan bool) {
	defer wg.Done()

	for {

	}
}
