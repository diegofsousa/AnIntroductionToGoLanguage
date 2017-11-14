package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println("Sevidor rodando...")
	rotas := mux.NewRouter().StrictSlash(true)
	rotas.HandleFunc("/membros", getMembros).Methods("GET")
	rotas.HandleFunc("/membros", postMembros).Methods("POST")
	rotas.HandleFunc("/membros/{id}", getMembro).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", rotas))

}

type Membro struct {
	ID   uint64 `json:"id"`
	Nome string `json:"nome"`
}

var membros = []Membro{
	Membro{ID: 1, Nome: "Diego Fernando"},
	Membro{ID: 2, Nome: "Leo Silva"},
}

func getMembros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	json.NewEncoder(w).Encode(membros)
}

func postMembros(w http.ResponseWriter, r *http.Request) {
	var t Membro
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &t)
	fmt.Println(t)
	membros = append(membros, t)
	w.Header().Set("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(membros)
}

func getMembro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	vars := mux.Vars(r)
	id := vars["id"]
	v, _ := strconv.ParseUint(id, 10, 32)
	membro, err := getAuxMembro(v)
	if err != "" {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(membro)
	}
}

func getAuxMembro(id uint64) (Membro, string) {
	for key, _ := range membros {
		if id == membros[key].ID {
			return membros[key], ""
		}
	}
	return Membro{}, "Não encontrado"
}
