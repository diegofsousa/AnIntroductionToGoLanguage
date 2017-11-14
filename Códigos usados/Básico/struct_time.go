package main

import "fmt"

type TimeDoPiaui struct {
	nome         string
	n_vitorias   int
	n_derrotas   int
	classificado bool
}

type Times []TimeDoPiaui

func main() {
	times := Times{}
	times.AdiconarTime("River", 23, 7, true)
	times.AdiconarTime("Parnahyba Sport Club", 25, 5, true)
	times.AdiconarTime("SEP", 20, 10, false)
	times.MostrarTimes()
	times.RetirarTime(0)
	times.MostrarTimes()
}

func (t *Times) AdiconarTime(nome string, n_vitorias int, n_derrotas int, classificado bool) {
	novoTime := TimeDoPiaui{
		nome,
		n_vitorias,
		n_derrotas,
		classificado,
	}

	*t = append(*t, novoTime)
	fmt.Printf("\n%s adicionado!", nome)
}

func (t *Times) RetirarTime(indice int) {
	times := *t
	time_retirado := times[indice]
	*t = append(times[0:indice], times[indice+1:]...)
	fmt.Printf("\n%s removido!", time_retirado.nome)
}

func (t Times) MostrarTimes() {
	fmt.Println("\n\n**Times:**")
	for i := 0; i < len(t); i++ {
		fmt.Printf("%s, %d, %d, %t\n", t[i].nome, t[i].n_vitorias, t[i].n_derrotas, t[i].classificado)
	}
}
