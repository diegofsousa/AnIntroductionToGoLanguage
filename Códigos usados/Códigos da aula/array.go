package main

import "fmt"

func main() {
	var lista []int
	pares := []int{2, 4, 6, 7}
	pares = append(pares[1:], pares[:2]...)
	impares := []int{3, 5, 7}
	nomes := []string{}
	fmt.Println(lista, pares, impares, nomes)
}
