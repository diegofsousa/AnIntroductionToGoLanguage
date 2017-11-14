package main

import "fmt"

func main() {
	b := true

	if b {
		fmt.Println("IF")
	} else {
		fmt.Println("ELSE")
	}

	number := 7

	switch number {
	case 2:
		fmt.Println("Number 2")
	case 5:
		fmt.Println("Number 5")
	default:
		fmt.Println("default")
	}

	if m := 5; m > 0 {
		fmt.Println("maior que zero")
	}

}
