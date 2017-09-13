package main

import "fmt"

func main() {
	sl := make([]int, 2)
	sl[0] = 3
	sl[1] = 78

	sl2 := sl[:]

	sl = append(sl, 1, 60)

	sl2[0] = 200

	fmt.Println(sl)
	fmt.Println(sl2)
	fmt.Println(len(sl2))

}
