package main

import (
	"fmt"
	"github.com/diegofsousa/distanceToEnucomp"
)

func main() {
	fmt.Println("Descrição do Evento:")
	fmt.Println(distanceToEnucomp.Info())

	fmt.Println("\nMinicursos do Evento:")
	for _, value := range distanceToEnucomp.ShortCourses() {
		fmt.Println(value)
	}

	fmt.Printf("\nA distancia até o Evento é de %f Km.\n",
		distanceToEnucomp.DistanceTo(-7.079602, -41.433254))
}
