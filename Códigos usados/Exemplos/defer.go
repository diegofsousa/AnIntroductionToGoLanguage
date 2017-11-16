package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("Segunda ação")
	}()

	fmt.Println("Primeira ação")
}
