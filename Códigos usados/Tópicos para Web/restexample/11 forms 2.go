package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", myHandle)
	http.ListenAndServe(":8080", nil)
}

var val string

func myHandle(res http.ResponseWriter, req *http.Request) {
	key := "q"
	fmt.Println(texto)
	val = req.FormValue(key)
	fmt.Println("value: ", val)

	res.Header().Set("Content-Type", "text/html")
	io.WriteString(res, texto)
}

var texto = `
<form method="POST">
<input type="text" name="q">
<input type="submit">
</form>
` + val
