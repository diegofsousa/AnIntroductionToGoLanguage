package models

type Pessoa struct {
	Nome  string
	Idade int
}

var pessoas = []Pessoa{
	Pessoa{Nome: "Diego", Idade: 21},
	Pessoa{Nome: "Fernando", Idade: 23},
}

func (p Pessoa) GetAll() []Pessoa {
	return pessoas
}
