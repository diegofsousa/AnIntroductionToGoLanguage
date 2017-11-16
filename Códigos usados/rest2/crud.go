package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"rest2/models"
	"strconv"
)

func main() {
	/*e := models.Estudiante{
		Name:   "Diego",
		Age:    30,
		Active: true,
	}

	err := models.Crear(e)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Objeto criado com exito")
	*/

	/*
		es, err := models.Consultar()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(es)
	*/

	/*
		e := models.Estudiante{
			ID:     2,
			Name:   "Diego",
			Age:    21,
			Active: false,
		}

		err := models.Actualizar(e)

		if err != nil {
			log.Fatal(err)
		}
	*/

	/*
		err := models.Deletar(3)
		if err != nil {
			log.Fatal(err)
		}
	*/
	/*
		es, err := models.Pesquisa("difed")

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(es)
	*/
	fmt.Println("Sevidor Rodando...")
	rotas := mux.NewRouter().StrictSlash(true)
	rotas.HandleFunc("/estudante", getUsers).Methods("GET")
	rotas.HandleFunc("/estudante", getUsersPost).Methods("POST")
	rotas.HandleFunc("/estudante/{id}/", getUser).Methods("GET")
	rotas.HandleFunc("/estudante/{id}/", getUserUpdate).Methods("PUT")
	rotas.HandleFunc("/estudante/{id}/", userDelete).Methods("DELETE")
	rotas.HandleFunc("/estudante/search/q={termo}/", getUserSearch).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", rotas))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	es, err := models.Consultar()

	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(es)
	}

}

func getUsersPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("POST")
	var e models.Estudiante
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &e)
	fmt.Println(e)

	err := models.Crear(e)

	if err != nil {
		log.Fatal(err)
	}

	es, err := models.Consultar()

	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(es)
	}

}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]
	v, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println(err)
	}

	es, err := models.ConsultaParaUm(v)

	if err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(http.StatusText(404))
	} else {
		json.NewEncoder(w).Encode(es)
	}

}

func getUserUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]
	v, err := strconv.Atoi(id)

	if err != nil {
		log.Fatal(err)
	}

	es, err := models.ConsultaParaUm(v)

	if err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(http.StatusText(404))
	} else {
		var aux models.Estudiante
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &aux)
		if aux.Name != "" {
			es.Name = aux.Name
		}

		if aux.Age != 0 {
			es.Age = aux.Age
		}

		if aux.Active != es.Active {
			es.Active = aux.Active
		}

		err := models.Actualizar(es)

		if err != nil {
			log.Fatal(err)
		}

		enew, err := models.ConsultaParaUm(es.ID)

		if err != nil {
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(http.StatusText(404))
		} else {
			json.NewEncoder(w).Encode(enew)
		}

	}
}

func getUserSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	termo := vars["termo"]

	es, err := models.Pesquisa(termo)

	if err != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(http.StatusText(404))
	} else {
		json.NewEncoder(w).Encode(es)
	}

}

func userDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	v, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println(err)
	}
	err = models.Deletar(v)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	es, err := models.Consultar()

	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(es)
	}
}
