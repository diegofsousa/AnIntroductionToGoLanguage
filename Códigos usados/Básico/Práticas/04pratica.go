package main

import (
	"fmt"
)

func main() {
	i := Inteiros{
		4,
		7,
	}

	f := Floats{
		5.78,
		4.32,
	}
	Show(i)
	Show(f)

}

func Show(i Calcular) {
	i.Soma()
}

type Calcular interface {
	Soma()
}

type Inteiros struct {
	x int
	y int
}

type Floats struct {
	x float64
	y float64
}

func (i Inteiros) Soma() {
	fmt.Println("Soma de inteiros = ", i.x+i.y)
}
func (f Floats) Soma() {
	fmt.Println("Soma de inteiros = ", f.x+f.y)
}
