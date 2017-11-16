package main

import (
	"fmt"
)

func main() {
	var idade int

	//idade = 21
	str := "Digite sua idade: "
	fmt.Print(str)
	fmt.Scan(&idade)
	fmt.Println("\nSua idade é: ", idade)
	if idade < 18 {
		fmt.Println("Menor que 18")
	} else {
		fmt.Println("Maior que 18")
	}
	//fmt.Println("hello world")
}
