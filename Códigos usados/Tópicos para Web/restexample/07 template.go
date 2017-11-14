package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", myHandleFunc)
	http.ListenAndServe(":8080", nil)
}

func myHandleFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	tmpl, err := template.New("anyNameForTemplate").Parse(doc)
	if err == nil {
		tmpl.Execute(w, nil)
	}
}

const doc = `
<!DOCTYPE html>
<html>
<head>
	<title>Primeiro template</title>
</head>
<body>
	<h1>Primeiro template</h1>
</body>
</html>

`
