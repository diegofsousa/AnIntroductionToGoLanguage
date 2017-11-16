package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func loopCount(name string) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		time.Sleep(1 + time.Second)
		fmt.Printf("Tarefa difícil -> %s\n", name)
	}
	fmt.Printf("Tarefa difícil %s acabou!\n", name)
}

func main() {
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go loopCount(strconv.Itoa(i))
	}

	wg.Wait()
	fmt.Println("Terminou")
}
