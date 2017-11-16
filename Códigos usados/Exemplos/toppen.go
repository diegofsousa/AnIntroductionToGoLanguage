package main

import "fmt"

func main() {
	idade := 22

	if idade <= 18 {
		fmt.Println("Menor ou igual a 18 anos")
	} else if idade > 25 {
		fmt.Println("Maior que 20 anos")
	} else {
		fmt.Println("Entre 19 e 25 anos")
	}
}
