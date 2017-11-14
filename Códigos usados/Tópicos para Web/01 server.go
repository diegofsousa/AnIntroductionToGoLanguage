package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Servidor rodando")
	http.HandleFunc("/", someFunc)
	http.ListenAndServe(":8080", nil)

}

func someFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello World"))
}
