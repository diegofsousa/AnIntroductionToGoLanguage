package main

import (
	"fmt"
	"net/http"
)

func main() {
	myMux := http.NewServeMux()
	fmt.Println("Servidor rodando")
	myMux.HandleFunc("/", someFunc)
	http.ListenAndServe(":8080", myMux)

}

func someFunc(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Request")
	w.Write([]byte("Hello World"))
}
