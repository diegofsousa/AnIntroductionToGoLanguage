package main

import (
	"fmt"
)

func main() {
	t := 2
	a := 5
	enucomp := "e"

	switch {
	case a == 7:
		fmt.Println("Entrou no 5")
	case t == 7:
		fmt.Println("Entrou no 2")
	case enucomp == "et":
		fmt.Println("Entrou em enucomp")
	default:
		fmt.Println("Nenhum dos casos")
	}

}
