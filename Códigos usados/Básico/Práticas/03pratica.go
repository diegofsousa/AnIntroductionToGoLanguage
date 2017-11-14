package main

import "fmt"

func main() {
	diego := Aluno{
		nome:   "Diego",
		idade:  21,
		status: "Estudando",
	}
	ana := Aluno{
		nome:   "Ana",
		idade:  19,
		status: "Estudando",
	}

	rodrigo := Professor{
		nome:      "Rodrigo",
		discipina: "Algebra Linear",
	}

	turma := Turma{}
	turma.alunos = append(turma.alunos, diego, ana)
	turma.professor = rodrigo
	fmt.Println(turma)
	turma.AprovarTodos()
	fmt.Println(turma)

}

type Aluno struct {
	nome   string
	idade  int
	status string
}

type Professor struct {
	nome      string
	discipina string
}

type Turma struct {
	alunos    []Aluno
	professor Professor
}

func (t *Turma) AprovarTodos() {
	for key, _ := range t.alunos {
		t.alunos[key].status = "Aprovado"
	}
}
