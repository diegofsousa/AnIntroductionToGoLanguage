package main

import (
	"net/http"
)

type Person struct {
	fName string
}

func (p *Person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("First Name: " + p.fName))
}

func main() {
	personOne := Person{fName: "Diego"}
	http.ListenAndServe(":8080", &personOne)
}
