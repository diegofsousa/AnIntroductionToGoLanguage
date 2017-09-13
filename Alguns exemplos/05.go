package main

import "fmt"

func main() {
	tabuada := [10]int{0, 5, 10} //array
	user := map[string]string{   //map
		"name": "Diego",
		"nick": "diegofsousa",
	}
	slice := []int{0, 5, 10} //slice

	tabuada[3] = 15

	slice = append(slice, 90, 110, 50)

	fmt.Println(len(tabuada))
	fmt.Println(tabuada)

	fmt.Println(user)
	fmt.Println(user["nick"])
	fmt.Println(slice)
}
