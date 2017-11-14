package main

import "fmt"

func main() {
	array := [10]int{}       //Array - Adicionável até 10
	array2 := []int{1, 2, 3} //Slice - Adicionável com append infinitas posições
	sl := make([]int, 2, 6)  //Slice com  2 posições e capacidade para 6

	sl[0] = 5
	sl[1] = 6
	sl = append(sl, 56)

	for i := 0; i < 10; i++ {
		array[i] = i
	}

	for i := 0; i < 3; i++ {
		array2[i] = i
	}

	fmt.Println(array)
	fmt.Println(array2)
	fmt.Println(sl)
}
