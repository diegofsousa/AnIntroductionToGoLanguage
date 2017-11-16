package main

import (
	"fmt"
	"sync"
	"time"
)

var group sync.WaitGroup

func sequencia(n int) {
	defer group.Done()
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", i+1)
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Printf(" ...Sequencia de %d terminou... ", n)
}

func main() {
	group.Add(3)
	go sequencia(10)
	go sequencia(15)
	go sequencia(20)
	group.Wait()
}
