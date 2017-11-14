package main

import "fmt"

type Jogador struct {
	nome    string
	numero  int
	posicao string
	overal  float64
}

type jogadores []Jogador

func main() {
	fmt.Println("...")
	j := jogadores{}

	fmt.Println(
		j.Adicionar(
			Jogador{
				"Diego",
				34,
				"Meia",
				56.4,
			}))
	fmt.Println(
		j.Adicionar(
			Jogador{
				"Diego",
				34,
				"Meia",
				56.4,
			}))

	fmt.Println(
		j.Adicionar(
			Jogador{
				"Dudu",
				34,
				"Meia",
				57.4,
			}))
	fmt.Println(j.Exibir())
	fmt.Println(j.Deletar(35))
	fmt.Println(j.Exibir())

}

func (j *jogadores) Adicionar(jogador Jogador) string {
	for _, value := range *j {
		if value.numero == jogador.numero {
			return "Número já existente"
		}
	}
	*j = append(*j, jogador)
	return "Adicionado!"
}

func (j jogadores) Exibir() jogadores {
	return j
}

func (j *jogadores) Deletar(numero int) string {
	l := *j
	for key, value := range *j {
		if value.numero == numero {
			*j = append(l[0:key], l[key+1:]...)
			return "Deletado"
		}
	}
	return "Número não existente"

}
