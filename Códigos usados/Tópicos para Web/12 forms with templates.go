package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"text/template"
)

type Person struct {
	FirstName string
	LastName  string
}

func main() {
	tpl, err := template.ParseFiles("templates/form.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fName := req.FormValue("first")
		lName := req.FormValue("last")
		fmt.Println("fName: ", fName)
		fmt.Println("[]byte(fName): ", []byte(fName))
		fmt.Println("typeOf: ", reflect.TypeOf(fName))

		err = tpl.Execute(res, Person{fName, lName})
		if err != nil {
			http.Error(res, err.Error(), 500)
			log.Fatal(err)
		}
	})

	http.ListenAndServe(":8080", nil)
}
