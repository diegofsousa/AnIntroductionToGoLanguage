package main

import (
	"fmt"
	"time"
)

func sequencia(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", i+1)
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Printf(" ...Sequencia de %d terminou... ", n)
}

func main() {
	go sequencia(10)
	go sequencia(10)
}
