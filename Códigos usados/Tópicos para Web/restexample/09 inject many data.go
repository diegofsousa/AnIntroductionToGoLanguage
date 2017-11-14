package main

import (
	"net/http"
	"text/template"
)

type Context struct {
	FirstName string
	Message   string
}

const doc = `
<!DOCTYPE html>
<html>
<head lang='pt-br'>
	<title>Primeiro template</title>
</head>
<body>
	<h1>Meu nome é {{ .FirstName }}</h1>
	<p>Recurso acessado: {{ .Message }}</p>
</body>
</html>
`

func toddFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	tmpl, err := template.New("anyNameForTemplate").Parse(doc)
	if err == nil {
		context := Context{
			"Tedd",
			"Mais Go, por favor",
		}
		tmpl.Execute(w, context)
	}
}

func main() {
	http.HandleFunc("/", toddFunc)
	http.ListenAndServe(":8080", nil)
}
