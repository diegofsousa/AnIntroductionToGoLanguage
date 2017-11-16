package main

import (
	"fmt"
	"strings"
)

func main() {
	//intervalo(10, 23)
	//numeros(1, 2, 3, 4, 5)
	maiusculo := func(str string) string {
		return strings.ToUpper(str)
	}

	nome := "Diego Fernando"

	fmt.Println(nome)
	fmt.Println(maiusculo(nome))
}

func imprimirString(nome string, idade int) {
	fmt.Printf("Olá, meu nome é %s e eu tenho %d anos.\n", nome, idade)
}

func intervalo(x, y int) {
	for i := x; i < y; i++ {
		fmt.Printf("%d ", i)
	}
}

func sucessorEAntecessor(x int) (int, int) {
	return x + 1, x - 1
}

func simplesSoma(x, y int) (soma int) {
	soma = x + y
	return
}

func numeros(lista ...int) {
	for _, numero := range lista {
		fmt.Println(numero)
	}
}
