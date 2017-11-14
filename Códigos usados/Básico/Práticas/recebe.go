package main

import "fmt"

func main() {
	var i, i2 int
	fmt.Printf("Informe os dois valores: ")
	fmt.Scanf("%d %d", &i, &i2)
	fmt.Println(i + i2)
}
