package main

import "fmt"

type TimeFutebol struct {
	n_vitorias int
}

type TimeBasquete struct {
	n_vitorias int
}

func (t TimeFutebol) PontosEmVitorias() int {
	return t.n_vitorias * 3
}

func (b TimeBasquete) PontosEmVitorias() int {
	return b.n_vitorias * 2
}

type Time interface {
	PontosEmVitorias() int
}

func Pontos(t Time) int {
	return t.PontosEmVitorias()
}

func main() {
	time_futebol := TimeFutebol{20}
	time_basquete := TimeBasquete{20}

	fmt.Println(Pontos(time_futebol))
	fmt.Println(Pontos(time_basquete))
}
