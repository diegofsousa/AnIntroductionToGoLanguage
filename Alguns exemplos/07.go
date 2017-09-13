package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3}

	user := map[string]string{
		"name": "diego",
		"nick": "leo",
	}

	sl := []string{
		"Go",
		"Python",
		"R",
	}

	for _, value := range arr {
		fmt.Println(value)
	}

	for key, value := range user {
		fmt.Printf("O campo %s tem o valor igual a %s\n", key, value)
	}

	for _, value := range sl {
		fmt.Println(value)
	}
}
